// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package traffic

import (
	"context"
	"time"

	"jhb-api/api/internal/svc"
	"jhb-api/api/internal/types"
	"jhb-api/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddTrafficLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 添加里程记录
func NewAddTrafficLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddTrafficLogic {
	return &AddTrafficLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddTrafficLogic) AddTraffic(req *types.AddTrafficReq) (resp *types.Response, err error) {
	// 验证用户是否存在
	_, err = l.svcCtx.SysUserModel.FindOne(l.ctx, req.UserId)
	if err != nil {
		if err == model.ErrNotFound {
			return &types.Response{
				Code:    404,
				Message: "用户不存在",
			}, nil
		}
		l.Errorf("Failed to get user: %v", err)
		return nil, err
	}

	// 解析时间
	trafficTime, err := time.Parse("2006-01-02 15:04:05", req.TrafficTime)
	if err != nil {
		return &types.Response{
			Code:    400,
			Message: "时间格式错误，请使用 yyyy-MM-dd HH:mm:ss 格式",
		}, nil
	}

	// 创建里程记录
	record := &model.TrafficRecord{
		UserId:      req.UserId,
		PlateNumber: req.PlateNumber,
		Mileage:     req.Amount, // 使用 Amount 作为 Mileage
		TrafficTime: trafficTime,
		IsCalculate: 0, // 默认未核算
	}

	_, err = l.svcCtx.TrafficRecordModel.Insert(l.ctx, record)
	if err != nil {
		l.Errorf("Failed to add traffic record: %v", err)
		return &types.Response{
			Code:    500,
			Message: "添加里程记录失败",
		}, nil
	}

	resp = &types.Response{
		Code:    0,
		Message: "添加成功",
	}

	return
}