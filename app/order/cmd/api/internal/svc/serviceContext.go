package svc

import (
	"go-bc/app/order/cmd/api/internal/config"
	"go-bc/app/order/cmd/rpc/order"
	"go-bc/app/payment/cmd/rpc/payment"
	"go-bc/app/travel/cmd/rpc/travel"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config

	OrderRpc   order.Order
	PaymentrPC payment.Payment
	TravelRpc  travel.Travel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,

		OrderRpc:   order.NewOrder(zrpc.MustNewClient(c.OrderRpcConf)),
		PaymentrPC: payment.NewPayment(zrpc.MustNewClient(c.PaymentRpcConf)),
		TravelRpc:  travel.NewTravel(zrpc.MustNewClient(c.TravelRpcConf)),
	}
}
