package auth

import (
	"errors"

	"github.com/Ghytro/inttest-configurator/internal/entity"
	"github.com/Ghytro/inttest-configurator/internal/entity/auth"
	"github.com/samber/lo"
)

func UserHasPerms(user auth.User, perms ...auth.EPermission) error {
	perms = lo.Uniq(perms)
	foundPermIds := lo.Uniq(
		lo.FlatMap(user.Roles, func(r auth.Role, _ int) []entity.BigIntPK {
			return lo.FilterMap(r.Perms, func(p auth.Perm, _ int) (entity.BigIntPK, bool) {
				return p.Id, lo.Contains(perms, p.Name)
			})
		}),
	)
	if len(foundPermIds) != len(perms) {
		return errors.New("недостаточно привелегий для совершения операции")
	}
	return nil
}
