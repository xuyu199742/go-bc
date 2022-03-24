package logic

import (
	"context"

	"go-bc/app/payment/cmd/rpc/internal/svc"
	"go-bc/app/payment/cmd/rpc/pb"
	"go-bc/app/payment/model"
	"go-bc/common/xerr"

	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetPaymentBySnLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPaymentBySnLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPaymentBySnLogic {
	return &GetPaymentBySnLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetPaymentBySnLogic) GetPaymentBySn(in *pb.GetPaymentBySnReq) (*pb.GetPaymentBySnResp, error) {

	thirdPayment, err := l.svcCtx.ThirdPaymentModel.FindOneBySn(in.Sn)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "查询第三方支付流水记录失败 db err:%v , in : %+v", err, in)
	}

	var resp pb.PaymentDetail
	if thirdPayment != nil {
		_ = copier.Copy(&resp, thirdPayment)

		resp.CreateTime = thirdPayment.CreateTime.Unix()
		resp.UpdateTime = thirdPayment.UpdateTime.Unix()
		resp.PayTime = thirdPayment.PayTime.Unix()
	}

	return &pb.GetPaymentBySnResp{
		PaymentDetail: &resp,
	}, nil
}
