package logic

import (
	"IMM_server/imm_user/user_api/internal/svc"
	"IMM_server/imm_user/user_api/internal/types"
	"IMM_server/imm_user/user_models"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type FriendListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFriendListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FriendListLogic {
	return &FriendListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FriendListLogic) FriendList(req *types.FriendListRequest) (resp *types.FriendListResponse, err error) {

	var count int64
	l.svcCtx.DB.Model(user_models.FriendModel{}).Where("send_user_id = ? or rev_user_id = ?", req.UserID, req.UserID).Count(&count)
	var friends []user_models.FriendModel
	l.svcCtx.DB.Preload("SendUserModel").Preload("RevUserModel").Find(&friends, "send_user_id = ? or rev_user_id = ?", req.UserID, req.UserID)
	var list []types.FriendInfoResponse
	for _, friend := range friends {

		info := types.FriendInfoResponse{}
		if friend.SendUserID == req.UserID {
			// 我是发起方
			info = types.FriendInfoResponse{
				UserID:   friend.RevUserID,
				Nickname: friend.RevUserModel.Nickname,
				Abstract: friend.RevUserModel.Abstract,
				Avatar:   friend.RevUserModel.Avatar,
				Notice:   friend.RevUserNotice,
			}
		}
		if friend.RevUserID == req.UserID {
			// 我是接收方
			info = types.FriendInfoResponse{
				UserID:   friend.SendUserID,
				Nickname: friend.SendUserModel.Nickname,
				Abstract: friend.SendUserModel.Abstract,
				Avatar:   friend.SendUserModel.Avatar,
				Notice:   friend.SenUserNotice,
			}
		}
		list = append(list, info)
	}

	return &types.FriendListResponse{List: list, Count: int(count)}, nil
}
