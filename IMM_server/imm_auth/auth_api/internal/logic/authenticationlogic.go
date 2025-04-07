package logic

import (
	"IMM_server/imm_auth/auth_api/internal/svc"
	"IMM_server/imm_auth/auth_api/internal/types"
	"IMM_server/utils"
	"IMM_server/utils/jwts"
	"context"
	"errors"
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"
)

type AuthenticationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAuthenticationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AuthenticationLogic {
	return &AuthenticationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AuthenticationLogic) Authentication(req *types.AuthenticationRequest) (resp string, err error) {

	fmt.Println(req.Token, req.ValidPath)
	if utils.InList(l.svcCtx.Config.WhiteList, req.ValidPath) {
		logx.Infof("%s 在白名单中", req.ValidPath)
		return "ok", nil
	}

	if req.Token == "" {
		err = errors.New("认证失败")
		return
	}

	payload, err := jwts.ParseToken(req.Token, l.svcCtx.Config.Auth.AccessSecret)
	if err != nil {
		err = errors.New("认证失败")
		return
	}

	_, err = l.svcCtx.Redis.Get(fmt.Sprintf("logout_%d", payload.UserID)).Result()
	if err == nil {
		err = errors.New("认证失败")
		return
	}
	return "ok", nil
}
