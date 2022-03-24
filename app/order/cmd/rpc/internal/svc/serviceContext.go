package svc

import (
	"go-bc/app/mqueue/cmd/rpc/mqueue"
	"go-bc/app/order/cmd/rpc/internal/config"
	"go-bc/app/order/model"
	"go-bc/app/travel/cmd/rpc/travel"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config

	TravelRpc travel.Travel
	MqueueRpc mqueue.Mqueue

	HomestayOrderModel model.HomestayOrderModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,

		TravelRpc: travel.NewTravel(zrpc.MustNewClient(c.TravelRpcConf)),
		MqueueRpc: mqueue.NewMqueue(zrpc.MustNewClient(c.MqueueRpcConf)),

		HomestayOrderModel: model.NewHomestayOrderModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
	}
}
