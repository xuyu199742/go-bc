package svc

import (
	"go-bc/app/travel/cmd/rpc/internal/config"
	"go-bc/app/travel/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config

	HomestayModel model.HomestayModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,

		HomestayModel: model.NewHomestayModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
	}
}
