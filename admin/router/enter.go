package router

import (
	"go-bc/admin/router/autocode"
	"go-bc/admin/router/banner"
	"go-bc/admin/router/example"
	"go-bc/admin/router/system"
)

type RouterGroup struct {
	System   system.RouterGroup
	Example  example.RouterGroup
	Autocode autocode.RouterGroup
	Banner   banner.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
