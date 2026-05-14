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

type GetUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取用户信息
func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserInfoLogic) GetUserInfo(req *types.UserIdReq) (resp *types.UserInfoResp, err error) {
	user, err := l.svcCtx.SysUserModel.FindOne(l.ctx, req.UserId)
	if err != nil {
		if err == model.ErrNotFound {
			l.Errorf("User not found: %d", req.UserId)
			return &types.UserInfoResp{
				Response: types.Response{
					Code:    404,
					Message: "用户不存在",
				},
			}, nil
		}
		l.Errorf("Failed to get user info: %v", err)
		return nil, err
	}

	resp = &types.UserInfoResp{
		Response: types.Response{
			Code:    0,
			Message: "success",
		},
		UserInfo: types.UserInfo{
			Id:          user.Id,
			UserName:    user.UserName,
			Phone:       user.Phone,
			PlateNumber: user.PlateNumber,
			CreateTime:  user.CreateTime.Format("2006-01-02 15:04:05"),
		},
	}

	return
}