package logic

import (
	"IMM_server/imm_auth/auth_api/internal/svc"
	"IMM_server/imm_auth/auth_api/internal/types"
	"IMM_server/imm_auth/auth_models"
	"IMM_server/imm_user/user_rpc/types/user_rpc"
	"IMM_server/utils/jwts"
	"IMM_server/utils/open_login"
	"context"
	"errors"

	"fmt"

	"github.com/zeromicro/go-zero/core/logx"
)

type Open_loginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOpen_loginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Open_loginLogic {
	return &Open_loginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Open_loginLogic) Open_login(req *types.OpenLoginRequest) (resp *types.LoginResponse, err error) {

	switch req.Flag {
	case "qq":
		info, err := open_login.NewQQLogin(req.Code, open_login.QQConfig{
			AppID:    l.svcCtx.Config.QQ.AppID,
			AppKey:   l.svcCtx.Config.QQ.AppKey,
			Redirect: l.svcCtx.Config.QQ.Redirect,
		})
		if err != nil {
			logx.Error(err)
			return nil, errors.New("登录失败")
		}
		fmt.Println(info)
		var user auth_models.UserModel
		err = l.svcCtx.DB.Take(&user, "open_id = ?", info.OpenID).Error
		if err != nil {
			// 注册逻辑
			fmt.Println("注册服务")
			res, err := l.svcCtx.UserRpc.UserCreate(context.Background(), &user_rpc.UserCreateRequest{
				NickName: info.Nickname,
				Password: "",
				Role:     2,
				Avatar:   info.Avatar,
				OpenId:   info.OpenID,
			})
			if err != nil {
				logx.Error(err)
				return nil, errors.New("登录失败")
			}
			user.Model.ID = uint(res.UserId)
			user.Role = 2
			user.Nickname = info.Nickname
		}

		// 登录逻辑
		token, err1 := jwts.GenToken(jwts.JwtPayLoad{
			UserID:   user.ID,
			Nickname: user.Nickname,
			Role:     user.Role,
		}, l.svcCtx.Config.Auth.AccessSecret, l.svcCtx.Config.Auth.AccessExpire)
		if err1 != nil {
			logx.Error(err1)
			err = errors.New("服务内部错误")
			return nil, err1
		}
		return &types.LoginResponse{Token: token}, nil
	}

	return
}
