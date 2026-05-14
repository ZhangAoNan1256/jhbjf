// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package point

import (
	"context"
	"fmt"

	"jhb-api/api/internal/svc"
	"jhb-api/api/internal/types"
	"jhb-api/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPointLogsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取积分日志
func NewGetPointLogsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPointLogsLogic {
	return &GetPointLogsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPointLogsLogic) GetPointLogs(req *types.UserIdReq) (resp *types.PointLogsResp, err error) {
	// 查询用户的积分日志
	query := fmt.Sprintf("select * from point_log where user_id = ? order by create_time desc")
	var logs []*model.PointLog
	err = l.svcCtx.SqlConn.QueryRowsCtx(l.ctx, &logs, query, req.UserId)
	if err != nil {
		l.Errorf("Failed to get point logs: %v", err)
		return nil, err
	}

	var list []types.PointLog
	for _, log := range logs {
		list = append(list, types.PointLog{
			Id:          log.Id,
			UserId:      log.UserId,
			ChangePoint: log.ChangePoint,
			ChangeType:  log.ChangeType,
			RelationId:  log.RelationId,
			CreateTime:  log.CreateTime.Format("2006-01-02 15:04:05"),
		})
	}

	resp = &types.PointLogsResp{
		Response: types.Response{
			Code:    0,
			Message: "success",
		},
		List: list,
	}

	return
}