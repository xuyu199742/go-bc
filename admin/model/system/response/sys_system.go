package response

import "go-bc/admin/config"

type SysConfigResponse struct {
	Config config.Server `json:"config"`
}
