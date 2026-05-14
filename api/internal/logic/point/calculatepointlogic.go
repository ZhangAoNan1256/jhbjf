// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package point

import (
	"context"

	"jhb-api/api/internal/svc"
	"jhb-api/api/internal/types"
	"jhb-api/model"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type CalculatePointLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 积分核算
func NewCalculatePointLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CalculatePointLogic {
	return &CalculatePointLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CalculatePointLogic) CalculatePoint(req *types.CalculatePointReq) (resp *types.CalculatePointResp, err error) {
	// 获取未核算的里程记录
	records, err := l.svcCtx.TrafficRecordModel.FindByUserId(l.ctx, req.UserId)
	if err != nil {
		l.Errorf("Failed to get traffic records: %v", err)
		return nil, err
	}

	// 筛选未核算的记录
	var uncalculatedRecords []*model.TrafficRecord
	totalAmount := 0.0
	for _, record := range records {
		if record.IsCalculate == 0 {
			uncalculatedRecords = append(uncalculatedRecords, record)
			totalAmount += record.Mileage
		}
	}

	if len(uncalculatedRecords) == 0 {
		return &types.CalculatePointResp{
			Response: types.Response{
				Code:    0,
				Message: "没有待核算的里程记录",
			},
			CalculatedPoints: 0,
		}, nil
	}

	// 获取积分规则
	rule, err := l.svcCtx.PointRuleModel.GetDefaultRule(l.ctx)
	if err != nil {
		l.Errorf("Failed to get point rule: %v", err)
		return nil, err
	}

	// 计算积分：消费金额 / 每元所需金额 * 积分值
	calculatedPoints := int(totalAmount / float64(rule.MilePerPoint) * float64(rule.PointValue))

	// 使用事务更新数据
	err = l.svcCtx.SqlConn.Transact(func(session sqlx.Session) error {
		ctx := l.ctx
		// 1. 更新里程记录状态为已核算
		for _, record := range uncalculatedRecords {
			record.IsCalculate = 1
			err := l.svcCtx.TrafficRecordModel.Update(ctx, record)
			if err != nil {
				return err
			}
		}

		// 2. 更新或创建用户积分
		userPoint, err := l.svcCtx.UserPointModel.FindOneByUserId(ctx, req.UserId)
		if err != nil {
			if err == model.ErrNotFound {
				// 创建新的积分记录
				userPoint = &model.UserPoint{
					UserId:     req.UserId,
					TotalPoint: int64(calculatedPoints),
				}
				_, err = l.svcCtx.UserPointModel.Insert(ctx, userPoint)
				if err != nil {
					return err
				}
			} else {
				return err
			}
		} else {
			// 更新积分余额
			userPoint.TotalPoint += int64(calculatedPoints)
			err = l.svcCtx.UserPointModel.Update(ctx, userPoint)
			if err != nil {
				return err
			}
		}

		// 3. 插入积分日志
		pointLog := &model.PointLog{
			UserId:      req.UserId,
			ChangePoint: int64(calculatedPoints),
			ChangeType:  "里程入账",
			RelationId:  0, // 关联多个里程记录，这里设为0
		}
		_, err = l.svcCtx.PointLogModel.Insert(ctx, pointLog)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		l.Errorf("Failed to calculate points: %v", err)
		return &types.CalculatePointResp{
			Response: types.Response{
				Code:    500,
				Message: "积分核算失败",
			},
			CalculatedPoints: 0,
		}, nil
	}

	resp = &types.CalculatePointResp{
		Response: types.Response{
			Code:    0,
			Message: "积分核算成功",
		},
		CalculatedPoints: calculatedPoints,
	}

	return
}