package listen

import (
	"context"
	"go-bc/app/message/cmd/mq/internal/config"
	"go-bc/app/message/cmd/mq/internal/mqs/kq/wxMessage"
	"go-bc/app/message/cmd/mq/internal/svc"

	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/service"
)

// kq
// 消息队列
func KqMqs(c config.Config, ctx context.Context, svcContext *svc.ServiceContext) []service.Service {

	return []service.Service{
		//监听消费流水状态变更.
		kq.MustNewQueue(c.SendWxMiniTplMessageConf, wxMessage.NewSendWxMiniSubMessageMq(ctx, svcContext)),
		//.....
	}

}
