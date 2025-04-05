package logic

import (
	"IMM_server/utils/jwts"
	"context"
	"errors"
	"fmt"

	"IMM_server/imm_auth/auth_api/internal/svc"
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

func (l *AuthenticationLogic) Authentication(token string) (resp string, err error) {
	if token == "" { // 空 Token 校验
		err = errors.New("认证失败")
		return
	}

	payload, err := jwts.ParseToken(token, l.svcCtx.Config.Auth.AccessSecret) // 使用 jwts.ParseToken 解密 Token
	if err != nil {
		err = errors.New("认证失败")
		return
	}

	_, err = l.svcCtx.Redis.Get(fmt.Sprintf("logout_%d", payload.UserID)).Result() // 如果用户已登出，则 Redis 会保存 logout_用户ID 的 key
	if err == nil {
		err = errors.New("认证失败")
		return
	}

	return "ok", nil
}
