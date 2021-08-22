package repository

import (
	"context"
	"github.com/danClauz/bibit/bmovie/search/shared"
	"github.com/danClauz/bibit/bmovie/search/shared/utils"
)

//go:generate mockgen -source=repository.go -destination=mocks/repository_mock.go -package=repository_mock Repository
type (
	Repository interface{
		Store(ctx context.Context, reqId string, obj interface{}) error
	}
	repo       struct {
		sh shared.Holder
	}
)

func New(sh shared.Holder) Repository {
	return &repo{sh: sh}
}

func (r *repo) Store(ctx context.Context, reqId string, obj interface{}) error {
	_ = r.sh.Logger.WithFields(map[string]interface{}{
		utils.XRequestId: reqId,
	})

	return nil
	//
	//if err := r.sh.Mysql.WithContext(ctx).Debug().Save(obj).Error; err != nil {
	//	logger.Errorf(logtag.ErrorTmpl, logtag.ErrRepositoryStore, err)
	//	return errors.Wrapf(err, "failed to store data %v", obj)
	//}
	//
	//return nil
}
