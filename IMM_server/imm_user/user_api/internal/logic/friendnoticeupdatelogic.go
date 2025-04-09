package logic

import (
	"IMM_server/imm_user/user_api/internal/svc"
	"IMM_server/imm_user/user_api/internal/types"
	"IMM_server/imm_user/user_models"
	"context"
	"errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type FriendNoticeUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFriendNoticeUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FriendNoticeUpdateLogic {
	return &FriendNoticeUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FriendNoticeUpdateLogic) FriendNoticeUpdate(req *types.FriendNoticeUpdateRequest) (resp *types.FriendNoticeUpdateResponse, err error) {

	var friend user_models.FriendModel

	if !friend.IsFriend(l.svcCtx.DB, req.UserID, req.FriendID) {
		return nil, errors.New("他还不是你的好友呢~")
	}
	if friend.SendUserID == req.UserID {
		// 我是发起方
		if friend.SenUserNotice == req.Notice {
			return
		}
		l.svcCtx.DB.Model(&friend).Update("sen_user_notice", req.Notice)
	}

	if friend.RevUserID == req.UserID {
		// 我是接收方
		if friend.RevUserNotice == req.Notice {
			return
		}
		l.svcCtx.DB.Model(&friend).Update("rev_user_notice", req.Notice)
	}

	return
}
