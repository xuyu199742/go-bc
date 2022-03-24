package homestay

import (
	"context"
	"go-bc/app/travel/cmd/api/internal/svc"
	"go-bc/app/travel/cmd/api/internal/types"
	"go-bc/app/travel/cmd/rpc/travel"
	"go-bc/common/tool"
	"go-bc/common/xerr"

	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type HomestayDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHomestayDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) HomestayDetailLogic {
	return HomestayDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HomestayDetailLogic) HomestayDetail(req types.HomestayDetailReq) (*types.HomestayDetailResp, error) {

	homestayResp, err := l.svcCtx.TravelRpc.HomestayDetail(l.ctx, &travel.HomestayDetailReq{
		Id: req.Id,
	})
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("获取民宿详情信息失败"), " id : %d , err : %v ", req.Id, err)
	}

	var typeHomestay types.Homestay
	if homestayResp.Homestay != nil {

		// 整合民宿详情
		_ = copier.Copy(&typeHomestay, homestayResp.Homestay)

		typeHomestay.FoodPrice = tool.Fen2Yuan(homestayResp.Homestay.FoodPrice)
		typeHomestay.HomestayPrice = tool.Fen2Yuan(homestayResp.Homestay.HomestayPrice)
		typeHomestay.MarketHomestayPrice = tool.Fen2Yuan(homestayResp.Homestay.MarketHomestayPrice)

	}

	return &types.HomestayDetailResp{
		Homestay: typeHomestay,
	}, nil
}
