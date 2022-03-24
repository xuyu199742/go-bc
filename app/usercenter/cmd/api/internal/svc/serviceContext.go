package svc

import (
	"go-bc/app/identity/cmd/rpc/identity"
	"go-bc/app/usercenter/cmd/api/internal/config"
	"go-bc/app/usercenter/cmd/rpc/usercenter"

	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config        config.Config
	UsercenterRpc usercenter.Usercenter
	IdentityRpc   identity.Identity

	SetUidToCtxMiddleware rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		UsercenterRpc: usercenter.NewUsercenter(zrpc.MustNewClient(c.UsercenterRpcConf)),
		IdentityRpc:   identity.NewIdentity(zrpc.MustNewClient(c.IdentityRpcConf)),
	}
}
