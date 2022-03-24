package svc

import (
	"go-bc/app/travel/cmd/api/internal/config"
	"go-bc/app/travel/cmd/rpc/travel"
	"go-bc/app/travel/model"
	"go-bc/app/usercenter/cmd/rpc/usercenter"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	//local
	Config config.Config

	//rpc
	UsercenterRpc usercenter.Usercenter
	TravelRpc     travel.Travel

	//model
	HomestayModel         model.HomestayModel
	HomestayActivityModel model.HomestayActivityModel
	HomestayBusinessModel model.HomestayBusinessModel
	HomestayCommentModel  model.HomestayCommentModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,

		UsercenterRpc: usercenter.NewUsercenter(zrpc.MustNewClient(c.UsercenterRpcConf)),
		TravelRpc:     travel.NewTravel(zrpc.MustNewClient(c.TravelRpcConf)),

		HomestayModel:         model.NewHomestayModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
		HomestayActivityModel: model.NewHomestayActivityModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
		HomestayBusinessModel: model.NewHomestayBusinessModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
		HomestayCommentModel:  model.NewHomestayCommentModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
	}
}
