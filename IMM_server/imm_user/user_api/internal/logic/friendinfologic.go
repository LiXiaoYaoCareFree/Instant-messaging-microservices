package logic

import (
	"IMM_server/imm_user/user_api/internal/svc"
	"IMM_server/imm_user/user_api/internal/types"
	"IMM_server/imm_user/user_models"
	"IMM_server/imm_user/user_rpc/types/user_rpc"
	"context"
	"encoding/json"
	"errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type FriendInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFriendInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FriendInfoLogic {
	return &FriendInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FriendInfoLogic) FriendInfo(req *types.FriendInfoRequest) (resp *types.FriendInfoResponse, err error) {
	// 确定你查的这个用户是自己的好友
	var friend user_models.FriendModel
	if !friend.IsFriend(l.svcCtx.DB, req.UserID, req.FriendID) {
		return nil, errors.New("他不是你的好友哦~")
	}
	res, err := l.svcCtx.UserRpc.UserInfo(context.Background(), &user_rpc.UserInfoRequest{
		UserId: uint32(req.FriendID),
	})
	if err != nil {
		return nil, errors.New(err.Error())
	}

	var friendUser user_models.UserModel
	json.Unmarshal(res.Data, &friendUser)

	response := types.FriendInfoResponse{
		UserID:   friendUser.ID,
		Nickname: friendUser.Nickname,
		Abstract: friendUser.Abstract,
		Avatar:   friendUser.Avatar,
		Notice:   friend.GetUserNotice(req.UserID),
	}

	return &response, nil
}
