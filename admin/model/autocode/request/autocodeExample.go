// 自动生成模板SysDictionaryDetail
package request

import (
	"go-bc/admin/model/autocode"
	"go-bc/admin/model/common/request"
)

// 如果含有time.Time 请自行import time包
type AutoCodeExampleSearch struct {
	autocode.AutoCodeExample
	request.PageInfo
}
