package logic

import (
	"IMM_server/common/models/ctype"
	"IMM_server/imm_user/user_api/internal/svc"
	"IMM_server/imm_user/user_api/internal/types"
	"IMM_server/imm_user/user_models"
	"IMM_server/utils/maps"
	"context"
	"errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoUpdateLogic {
	return &UserInfoUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoUpdateLogic) UserInfoUpdate(req *types.UserInfoUpdateRequest) (resp *types.UserInfoUpdateResponse, err error) {

	userMaps := maps.RefToMap(*req, "user")
	if len(userMaps) != 0 {
		var user user_models.UserModel
		err = l.svcCtx.DB.Take(&user, req.UserID).Error
		if err != nil {
			return nil, errors.New("用户不存在")
		}

		err = l.svcCtx.DB.Model(&user).Updates(userMaps).Error
		if err != nil {
			logx.Error(userMaps)
			logx.Error(err)
			return nil, errors.New("用户信息更新失败")
		}
	}
	userConfMaps := maps.RefToMap(*req, "user_conf")
	if len(userConfMaps) != 0 {
		var userConf user_models.UserConfModel
		err = l.svcCtx.DB.Take(&userConf, "user_id = ?", req.UserID).Error
		if err != nil {
			return nil, errors.New("用户配置不存在")
		}

		verificationQuestion, ok := userConfMaps["verification_question"]
		if ok {
			delete(userConfMaps, "verification_question")

			data := ctype.VerificationQuestion{}
			maps.MapToStruct(verificationQuestion.(map[string]any), &data)

			l.svcCtx.DB.Model(&userConf).Updates(&user_models.UserConfModel{
				VerificationQuestion: &data,
			})
		}

		err = l.svcCtx.DB.Model(&userConf).Updates(userConfMaps).Error

		if err != nil {
			logx.Error(userConfMaps)
			logx.Error(err)
			return nil, errors.New("用户信息更新失败")
		}
	}
	return
}
