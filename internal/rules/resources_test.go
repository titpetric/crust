package rules_test

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/titpetric/factory"

	"github.com/crusttech/crust/internal/auth"
	"github.com/crusttech/crust/internal/rules"
	"github.com/crusttech/crust/system/types"

	. "github.com/crusttech/crust/internal/test"
)

func TestRules(t *testing.T) {
	user := &types.User{ID: 1337}
	ctx := auth.SetIdentityToContext(context.Background(), user)

	db := factory.Database.MustGet()

	Error(t, db.Transaction(func() error {
		db.Insert("sys_user", user)
		var i uint64 = 0
		for i < 5 {
			db.Insert("sys_team", types.Team{ID: i, Name: fmt.Sprintf("Team %d", i)})
			i++
		}
		db.Insert("sys_team_member", types.TeamMember{TeamID: 1, UserID: user.ID})
		db.Insert("sys_team_member", types.TeamMember{TeamID: 2, UserID: user.ID})

		Expect := func(expected rules.Access, actual rules.Access, format string, params ...interface{}) {
			Assert(t, expected == actual, format, params...)
		}

		resources := rules.NewResources(ctx, db)

		// default (unset=deny)
		{
			Expect(rules.Inherit, resources.IsAllowed("channel:1", "edit"), "expected inherit")
			Expect(rules.Inherit, resources.IsAllowed("channel:*", "edit"), "expected inherit")
		}

		// allow channel:2 group:2 (default deny, multi=allow)
		{
			resources.Grant(2, "channel:2", []string{"edit", "delete"}, rules.Allow)
			Expect(rules.Inherit, resources.IsAllowed("channel:1", "edit"), "expected error, got nil")
			Expect(rules.Allow, resources.IsAllowed("channel:2", "edit"), "channel:2 edit, expected no error")
			Expect(rules.Allow, resources.IsAllowed("channel:*", "edit"), "channel:* edit, expected no error")
		}

		// list grants for team
		{
			grants, err := resources.ListGrants(2, "channel:2")
			NoError(t, err, "expect no error")
			Assert(t, len(grants) == 2, "expected 2 grants")
			Assert(t, grants[0].TeamID == 2, "expected TeamID == 2, got %v", grants[0].TeamID)
			Assert(t, grants[0].Resource == "channel:2", "expected Resource == channel:2, got %s", grants[0].Resource)
			Assert(t, grants[0].Operation == "delete", "expected Operation == delete, got %s", grants[0].Operation)
			Assert(t, grants[0].Value == rules.Allow, "expected Value == Allow, got %s", grants[0].Value)
		}

		// deny channel:1 group:1 (explicit deny, multi=deny)
		{
			resources.Grant(1, "channel:1", []string{"edit"}, rules.Deny)
			Expect(rules.Deny, resources.IsAllowed("channel:1", "edit"), "expected error, got nil")
			Expect(rules.Allow, resources.IsAllowed("channel:2", "edit"), "channel:2 edit, expected no error")
			Expect(rules.Deny, resources.IsAllowed("channel:*", "edit"), "expected error, got nil")
		}

		// reset (unset=deny)
		{
			resources.Grant(2, "channel:2", []string{"edit", "delete"}, rules.Inherit)
			resources.Grant(1, "channel:1", []string{"edit", "delete"}, rules.Inherit)
			Expect(rules.Inherit, resources.IsAllowed("channel:1", "edit"), "expected error, got nil")
			Expect(rules.Inherit, resources.IsAllowed("channel:*", "edit"), "expected error, got nil")
		}
		return errors.New("Rollback")
	}), "Expected rollback error, got nil")
}
