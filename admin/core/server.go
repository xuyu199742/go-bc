package core

import (
	"fmt"
	"go-bc/admin/rpc"
	"time"

	"go-bc/admin/global"
	"go-bc/admin/initialize"
	"go-bc/admin/service/system"
	"go.uber.org/zap"
)

type server interface {
	ListenAndServe() error
}

func RunWindowsServer() {

	//// log、prometheus、trace、metricsUrl.
	//if err := global.GVA_CONFIG.ServiceConf.SetUp(); err != nil {
	//	panic(err)
	//}

	if global.GVA_CONFIG.System.UseMultipoint {
		// 初始化redis服务
		initialize.Redis()
	}

	// 从db加载jwt数据
	if global.GVA_DB != nil {
		system.LoadAll()
	}

	//初始化rpc client
	rpc.InitClient()

	Router := initialize.Routers()

	Router.Static("/form-generator", "./resource/page")

	address := fmt.Sprintf(":%d", global.GVA_CONFIG.System.Addr)
	s := initServer(address, Router)
	// 保证文本顺序输出
	// In order to ensure that the text order output can be deleted
	time.Sleep(10 * time.Microsecond)
	global.GVA_LOG.Info("server run success on ", zap.String("address", address))

	global.GVA_LOG.Error(s.ListenAndServe().Error())
}
