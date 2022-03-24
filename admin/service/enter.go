package service

import (
	"go-bc/admin/service/autocode"
	"go-bc/admin/service/banner"
	"go-bc/admin/service/example"
	"go-bc/admin/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup   system.ServiceGroup
	ExampleServiceGroup  example.ServiceGroup
	AutoCodeServiceGroup autocode.ServiceGroup
	BannerServiceGroup   banner.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
