package projects

import (
	"configurator/internal/repository/internal/common"
	"configurator/pkg/database"

	"go.uber.org/zap"
)

type Repository struct {
	common.Mixin
}

func New(db *database.PGDB, log *zap.SugaredLogger) *Repository {
	return &Repository{
		Mixin: common.Mixin{
			DB:          db,
			Log:         log.With(zap.String("repo", "projects_repo")),
			ErrWrapDesc: "err in projects_repo",
		},
	}
}
