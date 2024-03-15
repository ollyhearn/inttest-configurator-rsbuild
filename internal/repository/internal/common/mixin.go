package common

import (
	"context"

	"github.com/Ghytro/inttest-configurator/pkg/database"
	"github.com/pkg/errors"
	"github.com/samber/lo"
	"go.uber.org/zap"
)

// fixme: public не выход (это го а не джава, nvm)

type Mixin struct {
	DB          *database.PGDB
	Log         *zap.SugaredLogger
	ErrWrapDesc string
}

func (r Mixin) RunInTransaction(ctx context.Context, fn func(ctx context.Context) error) error {
	ctx, err := r.DB.TxContext(ctx)
	if err != nil {
		return err
	}
	tx := lo.Must(database.GetPgTx(ctx))
	if err := fn(ctx); err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}

func (r Mixin) WrapErr(err error) error {
	return errors.Wrap(err, r.ErrWrapDesc)
}
