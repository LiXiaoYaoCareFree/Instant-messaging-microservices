package user_models

import (
	"IMM_server/common/models"
	"IMM_server/common/models/ctype"
)

// FriendVerifyModel 好友验证表
type FriendVerifyModel struct {
	models.Model
	SendUserID           uint                        `json:"sendUserID"`                         // 发起验证方
	SendUserModel        UserModel                   `gorm:"foreignKey:SendUserID" json:"-"`     // 发起验证方
	RevUserID            uint                        `json:"revUserID"`                          // 接受验证方
	RevUserModel         UserModel                   `gorm:"foreignKey:RevUserID" json:"-"`      // 接受验证方
	Status               int8                        `json:"status"`                             // 状态 0 未操作 1 同意 2 拒绝 3 忽略
	AdditionalMessages   string                      `gorm:"size:128" json:"additionalMessages"` // 附加消息
	VerificationQuestion *ctype.VerificationQuestion `json:"verificationQuestion"`               // 验证问题  为3和4的时候需要
}
