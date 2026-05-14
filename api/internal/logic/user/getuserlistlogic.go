// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package user

import (
	"context"

	"jhb-api/api/internal/svc"
	"jhb-api/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取用户列表
func NewGetUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserListLogic {
	return &GetUserListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserListLogic) GetUserList() (resp *types.UserListResp, err error) {
	users, err := l.svcCtx.SysUserModel.FindAll(l.ctx)
	if err != nil {
		l.Errorf("Failed to get user list: %v", err)
		return nil, err
	}

	var userList []types.UserInfo
	for _, user := range users {
		userList = append(userList, types.UserInfo{
			Id:          user.Id,
			UserName:    user.UserName,
			Phone:       user.Phone,
			PlateNumber: user.PlateNumber,
			CreateTime:  user.CreateTime.Format("2006-01-02 15:04:05"),
		})
	}

	resp = &types.UserListResp{
		Response: types.Response{
			Code:    0,
			Message: "success",
		},
		List: userList,
	}

	return
}