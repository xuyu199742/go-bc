package request

import (
	"go-bc/admin/model/autocode"
	"go-bc/admin/model/common/request"
)

type {{.StructName}}Search struct{
    autocode.{{.StructName}}
    request.PageInfo
}
