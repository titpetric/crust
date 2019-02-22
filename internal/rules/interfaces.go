package rules

import (
	"context"

	"github.com/titpetric/factory"
)

type ResourcesInterface interface {
	With(ctx context.Context, db *factory.DB) ResourcesInterface

	IsAllowed(resource string, operation string) Access

	Grant(roleID uint64, rules []Rule) error
	Read(roleID uint64) ([]Rule, error)
	Delete(roleID uint64) error
}
