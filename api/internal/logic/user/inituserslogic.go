// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package user

import (
	"context"

	"jhb-api/api/internal/svc"
	"jhb-api/api/internal/types"
	"jhb-api/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type InitUsersLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 初始化用户
func NewInitUsersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InitUsersLogic {
	return &InitUsersLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InitUsersLogic) InitUsers() (resp *types.Response, err error) {
	// 初始化示例用户数据
	users := []*model.SysUser{
		{
			UserName:    "张三",
			Phone:       "13800138001",
			PlateNumber: "京A88888",
		},
		{
			UserName:    "李四",
			Phone:       "13800138002",
			PlateNumber: "京B66666",
		},
		{
			UserName:    "王五",
			Phone:       "13800138003",
			PlateNumber: "京C99999",
		},
	}

	for _, user := range users {
		// 检查是否已存在
		existUser, _ := l.svcCtx.SysUserModel.FindOneByPhone(l.ctx, user.Phone)
		if existUser != nil {
			continue
		}

		_, err := l.svcCtx.SysUserModel.Insert(l.ctx, user)
		if err != nil {
			l.Errorf("Failed to insert user %s: %v", user.UserName, err)
			return &types.Response{
				Code:    500,
				Message: "初始化用户失败",
			}, nil
		}
	}

	resp = &types.Response{
		Code:    0,
		Message: "用户初始化成功",
	}

	return
}
