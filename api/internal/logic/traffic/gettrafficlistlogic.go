// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package traffic

import (
	"context"

	"jhb-api/api/internal/svc"
	"jhb-api/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTrafficListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取里程列表
func NewGetTrafficListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTrafficListLogic {
	return &GetTrafficListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTrafficListLogic) GetTrafficList(req *types.UserIdReq) (resp *types.TrafficListResp, err error) {
	records, err := l.svcCtx.TrafficRecordModel.FindByUserId(l.ctx, req.UserId)
	if err != nil {
		l.Errorf("Failed to get traffic list: %v", err)
		return nil, err
	}

	var list []types.TrafficRecord
	for _, record := range records {
		list = append(list, types.TrafficRecord{
			Id:          record.Id,
			UserId:      record.UserId,
			PlateNumber: record.PlateNumber,
			Amount:      record.Mileage, // 使用 Mileage 作为 Amount
			TrafficTime: record.TrafficTime.Format("2006-01-02 15:04:05"),
			IsCalculate: int(record.IsCalculate), // 转换类型
			CreateTime:  record.CreateTime.Format("2006-01-02 15:04:05"),
		})
	}

	resp = &types.TrafficListResp{
		Response: types.Response{
			Code:    0,
			Message: "success",
		},
		List: list,
	}

	return
}