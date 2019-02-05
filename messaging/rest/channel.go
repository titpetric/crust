package rest

import (
	"context"

	"github.com/crusttech/crust/internal/payload"
	"github.com/crusttech/crust/internal/payload/outgoing"
	"github.com/crusttech/crust/sam/rest/request"
	"github.com/crusttech/crust/sam/service"
	"github.com/crusttech/crust/sam/types"
	"github.com/pkg/errors"
)

var _ = errors.Wrap

type (
	Channel struct {
		svc struct {
			ch  service.ChannelService
			att service.AttachmentService
		}
	}
)

func (Channel) New() *Channel {
	ctrl := &Channel{}
	ctrl.svc.ch = service.DefaultChannel
	ctrl.svc.att = service.DefaultAttachment

	return ctrl
}

func (ctrl *Channel) Create(ctx context.Context, r *request.ChannelCreate) (interface{}, error) {
	channel := &types.Channel{
		Name:    r.Name,
		Topic:   r.Topic,
		Type:    types.ChannelType(r.Type),
		Members: payload.ParseUInt64s(r.Members),
	}

	return ctrl.wrap(ctrl.svc.ch.With(ctx).Create(channel))
}

func (ctrl *Channel) Update(ctx context.Context, r *request.ChannelUpdate) (interface{}, error) {
	channel := &types.Channel{
		ID:    r.ChannelID,
		Name:  r.Name,
		Topic: r.Topic,
		Type:  types.ChannelType(r.Type),
	}

	return ctrl.wrap(ctrl.svc.ch.With(ctx).Update(channel))
}

func (ctrl *Channel) State(ctx context.Context, r *request.ChannelState) (interface{}, error) {
	switch r.State {
	case "delete":
		return ctrl.wrap(ctrl.svc.ch.With(ctx).Delete(r.ChannelID))
	case "undelete":
		return ctrl.wrap(ctrl.svc.ch.With(ctx).Undelete(r.ChannelID))
	case "archive":
		return ctrl.wrap(ctrl.svc.ch.With(ctx).Archive(r.ChannelID))
	case "unarchive":
		return ctrl.wrap(ctrl.svc.ch.With(ctx).Unarchive(r.ChannelID))
	}

	return nil, nil
}

func (ctrl *Channel) SetFlag(ctx context.Context, r *request.ChannelSetFlag) (interface{}, error) {
	switch r.Flag {
	case "pinned", "hidden", "ignored":
		return ctrl.wrap(ctrl.svc.ch.With(ctx).SetFlag(r.ChannelID, types.ChannelMembershipFlag(r.Flag)))
	}

	return nil, nil
}

func (ctrl *Channel) RemoveFlag(ctx context.Context, r *request.ChannelRemoveFlag) (interface{}, error) {
	return ctrl.wrap(ctrl.svc.ch.With(ctx).SetFlag(r.ChannelID, types.ChannelMembershipFlagNone))
}

func (ctrl *Channel) Read(ctx context.Context, r *request.ChannelRead) (interface{}, error) {
	return ctrl.wrap(ctrl.svc.ch.With(ctx).FindByID(r.ChannelID))
}

func (ctrl *Channel) List(ctx context.Context, r *request.ChannelList) (interface{}, error) {
	return ctrl.wrapSet(ctrl.svc.ch.With(ctx).Find(&types.ChannelFilter{Query: r.Query}))
}

func (ctrl *Channel) Members(ctx context.Context, r *request.ChannelMembers) (interface{}, error) {
	return ctrl.wrapMemberSet(ctrl.svc.ch.With(ctx).FindMembers(r.ChannelID))
}

func (ctrl *Channel) Invite(ctx context.Context, r *request.ChannelInvite) (interface{}, error) {
	return ctrl.wrapMemberSet(ctrl.svc.ch.With(ctx).InviteUser(r.ChannelID, r.UserID...))
}

func (ctrl *Channel) Join(ctx context.Context, r *request.ChannelJoin) (interface{}, error) {
	return ctrl.wrapMemberSet(ctrl.svc.ch.With(ctx).AddMember(r.ChannelID, r.UserID))
}

func (ctrl *Channel) Part(ctx context.Context, r *request.ChannelPart) (interface{}, error) {
	return nil, ctrl.svc.ch.With(ctx).DeleteMember(r.ChannelID, r.UserID)
}

func (ctrl *Channel) Attach(ctx context.Context, r *request.ChannelAttach) (interface{}, error) {
	file, err := r.Upload.Open()
	if err != nil {
		return nil, err
	}

	defer file.Close()

	return ctrl.wrapAttachment(ctrl.svc.att.With(ctx).Create(
		r.Upload.Filename,
		r.Upload.Size,
		file,
		r.ChannelID,
		r.ReplyTo,
	))
}

func (ctrl *Channel) wrapAttachment(attachment *types.Attachment, err error) (*outgoing.Attachment, error) {
	if err != nil {
		return nil, err
	} else {
		return payload.Attachment(attachment), nil
	}
}

func (ctrl *Channel) wrap(channel *types.Channel, err error) (*outgoing.Channel, error) {
	if err != nil {
		return nil, err
	} else {
		return payload.Channel(channel), nil
	}
}

func (ctrl *Channel) wrapSet(cc types.ChannelSet, err error) (*outgoing.ChannelSet, error) {
	if err != nil {
		return nil, err
	} else {
		return payload.Channels(cc), nil
	}
}

func (ctrl *Channel) wrapMember(m *types.ChannelMember, err error) (*outgoing.ChannelMember, error) {
	if err != nil {
		return nil, err
	} else {
		return payload.ChannelMember(m), nil
	}
}

func (ctrl *Channel) wrapMemberSet(mm types.ChannelMemberSet, err error) (*outgoing.ChannelMemberSet, error) {
	if err != nil {
		return nil, err
	} else {
		return payload.ChannelMembers(mm), nil
	}
}