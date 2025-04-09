package logic

import (
	"IMM_server/common/list_query"
	"IMM_server/common/models"
	"IMM_server/imm_user/user_api/internal/svc"
	"IMM_server/imm_user/user_api/internal/types"
	"IMM_server/imm_user/user_models"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchLogic {
	return &SearchLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchLogic) Search(req *types.SearchRequest) (resp *types.SearchResponse, err error) {
	// 先找所有的用户
	friends, count, _ := list_query.ListQuery(l.svcCtx.DB, user_models.UserConfModel{}, list_query.Option{
		PageInfo: models.PageInfo{
			Page:  req.Page,
			Limit: req.Limit,
		},
		Preload: []string{"UserModel"},
		Where:   l.svcCtx.DB.Where("search_user <> 0 or (search_user = 1 and user_id = ?)", req.Key),
	})
	list := make([]types.SearchInfo, 0)
	for _, friend := range friends {
		list = append(list, types.SearchInfo{
			UserID:   friend.UserID,
			Nickname: friend.UserModel.Nickname,
			Abstract: friend.UserModel.Abstract,
			Avatar:   friend.UserModel.Avatar,
		})
	}

	return &types.SearchResponse{List: list, Count: count}, nil
}
