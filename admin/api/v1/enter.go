package v1

import (
	"go-bc/admin/api/v1/autocode"
	"go-bc/admin/api/v1/banner"
	"go-bc/admin/api/v1/example"
	"go-bc/admin/api/v1/system"
)

type ApiGroup struct {
	SystemApiGroup   system.ApiGroup
	ExampleApiGroup  example.ApiGroup
	AutoCodeApiGroup autocode.ApiGroup
	BannerApiGroup   banner.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
