package service

import (
	"context"

	"github.com/pkg/errors"
	"github.com/titpetric/factory"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/crusttech/crust/compose/internal/repository"
	"github.com/crusttech/crust/compose/types"
	"github.com/crusttech/crust/internal/logger"
)

type (
	page struct {
		db     *factory.DB
		ctx    context.Context
		logger *zap.Logger

		prmSvc PermissionsService

		pageRepo   repository.PageRepository
		moduleRepo repository.ModuleRepository
	}

	PageService interface {
		With(ctx context.Context) PageService

		FindByID(namespaceID, pageID uint64) (*types.Page, error)
		FindByModuleID(namespaceID, moduleID uint64) (*types.Page, error)
		FindBySelfID(namespaceID, selfID uint64) (pages types.PageSet, f types.PageFilter, err error)
		Find(filter types.PageFilter) (set types.PageSet, f types.PageFilter, err error)
		Tree(namespaceID uint64) (pages types.PageSet, err error)

		Create(page *types.Page) (*types.Page, error)
		Update(page *types.Page) (*types.Page, error)
		DeleteByID(namespaceID, pageID uint64) error

		Reorder(namespaceID, selfID uint64, pageIDs []uint64) error
	}
)

const (
	ErrModulePageExists serviceError = "ModulePageExists"
)

func Page() PageService {
	return (&page{
		logger: DefaultLogger.Named("page"),
		prmSvc: DefaultPermissions,
	}).With(context.Background())
}

func (svc page) With(ctx context.Context) PageService {
	db := repository.DB(ctx)
	return &page{
		db:     db,
		ctx:    ctx,
		logger: svc.logger,

		prmSvc: svc.prmSvc.With(ctx),

		pageRepo:   repository.Page(ctx, db),
		moduleRepo: repository.Module(ctx, db),
	}
}

// log() returns zap's logger with requestID from current context and fields.
func (svc page) log(fields ...zapcore.Field) *zap.Logger {
	return logger.AddRequestID(svc.ctx, svc.logger).With(fields...)
}

func (svc page) FindByID(namespaceID, pageID uint64) (p *types.Page, err error) {
	return svc.checkPermissions(svc.pageRepo.FindByID(namespaceID, pageID))
}

func (svc page) FindByModuleID(namespaceID, moduleID uint64) (p *types.Page, err error) {
	return svc.checkPermissions(svc.pageRepo.FindByModuleID(namespaceID, moduleID))
}

func (svc page) checkPermissions(p *types.Page, err error) (*types.Page, error) {
	if err != nil {
		return nil, err
	} else if !svc.prmSvc.CanReadPage(p) {
		return nil, errors.New("not allowed to access this page")
	}

	return p, err
}

func (svc page) FindBySelfID(namespaceID, parentID uint64) (pp types.PageSet, f types.PageFilter, err error) {
	if namespaceID == 0 {
		return nil, f, ErrNamespaceRequired.withStack()
	}

	return svc.filterPageSetByPermission(svc.pageRepo.Find(types.PageFilter{
		NamespaceID: namespaceID,
		ParentID:    parentID,

		// This will enable parentID=0 query
		Root: true,
	}))
}

func (svc page) Find(filter types.PageFilter) (set types.PageSet, f types.PageFilter, err error) {
	if filter.NamespaceID == 0 {
		return nil, f, ErrNamespaceRequired.withStack()
	}

	return svc.filterPageSetByPermission(svc.pageRepo.Find(filter))
}

func (svc page) Tree(namespaceID uint64) (pages types.PageSet, err error) {
	if namespaceID == 0 {
		return nil, ErrNamespaceRequired.withStack()
	}

	var (
		tree   types.PageSet
		filter = types.PageFilter{
			NamespaceID: namespaceID,
		}
	)

	return tree, svc.db.Transaction(func() (err error) {
		if pages, _, err = svc.filterPageSetByPermission(svc.pageRepo.Find(filter)); err != nil {
			return
		}

		// No preloading - we do not need (or should have) any modules
		// associated with us
		_ = pages.Walk(func(p *types.Page) error {
			if p.SelfID == 0 {
				tree = append(tree, p)
			} else if c := pages.FindByID(p.SelfID); c != nil {
				if c.Children == nil {
					c.Children = types.PageSet{}
				}

				c.Children = append(c.Children, p)
			} else {
				// Move orphans to root
				p.SelfID = 0
				tree = append(tree, p)
			}

			return nil
		})

		return nil
	})
}

func (svc page) filterPageSetByPermission(pp types.PageSet, f types.PageFilter, err error) (types.PageSet, types.PageFilter, error) {
	if err != nil {
		return nil, f, err
	}

	// @todo Filter-by-permission can/will mess up filter's count & paging...
	pp, err = pp.Filter(func(m *types.Page) (bool, error) {
		return svc.prmSvc.CanReadPage(m), nil
	})

	return pp, f, err
}

func (svc page) Reorder(namespaceID, selfID uint64, pageIDs []uint64) error {
	return svc.pageRepo.Reorder(namespaceID, selfID, pageIDs)
}

func (svc page) Create(mod *types.Page) (p *types.Page, err error) {
	mod.ID = 0

	if mod.NamespaceID == 0 {
		return nil, ErrNamespaceRequired.withStack()
	}

	if !svc.prmSvc.CanCreatePage(crmNamespace()) {
		return nil, ErrNoCreatePermissions.withStack()
	}

	if err = svc.checkModulePage(mod); err != nil {
		return
	}

	p, err = svc.pageRepo.Create(mod)
	return
}

func (svc page) Update(mod *types.Page) (p *types.Page, err error) {
	if mod.ID == 0 {
		return nil, ErrInvalidID.withStack()
	}

	if p, err = svc.pageRepo.FindByID(mod.NamespaceID, mod.ID); err != nil {
		return
	}

	if isStale(mod.UpdatedAt, p.UpdatedAt, p.CreatedAt) {
		return nil, ErrStaleData.withStack()
	}

	if !svc.prmSvc.CanUpdatePage(p) {
		return nil, ErrNoUpdatePermissions.withStack()
	}

	if err = svc.checkModulePage(mod); err != nil {
		return
	}

	p.ModuleID = mod.ModuleID
	p.SelfID = mod.SelfID
	p.Blocks = mod.Blocks
	p.Title = mod.Title
	p.Description = mod.Description
	p.Visible = mod.Visible
	p.Weight = mod.Weight

	p, err = svc.pageRepo.Update(p)
	return
}

func (svc page) checkModulePage(mod *types.Page) error {
	if mod.ModuleID > 0 {
		if p, err := svc.pageRepo.FindByModuleID(mod.NamespaceID, mod.ModuleID); err != nil {
			if err.Error() != repository.ErrPageNotFound.Error() {
				return err
			}
		} else if p.ID > 0 && mod.ID != p.ID {
			return ErrModulePageExists
		}
	}

	return nil
}

func (svc page) DeleteByID(namespaceID, pageID uint64) error {
	if p, err := svc.pageRepo.FindByID(namespaceID, pageID); err != nil {
		return errors.Wrap(err, "could not delete page")
	} else if !svc.prmSvc.CanDeletePage(p) {
		return errors.New("not allowed to delete this page")
	}

	return svc.pageRepo.DeleteByID(namespaceID, pageID)
}
