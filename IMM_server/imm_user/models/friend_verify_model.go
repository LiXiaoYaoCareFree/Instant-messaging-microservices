package models

import "IMM_server/common/models"

// FriendVerifyModel 好友验证表
type FriendVerifyModel struct {
	models.Model
	SendUserID           uint                  `json:"sendUserID"`           // 发起验证方
	RevUserID            uint                  `json:"revUserID"`            // 接收验证方
	Status               int8                  `json:"status"`               // 状态 0 未操作 1 同意 2 拒绝 3 忽略
	AdditionalMessages   string                `json:"additionalMessages"`   // 附加信息
	VerificationQuestion *VerificationQuestion `json:"verificationQuestion"` // 验证问题 为3和4的时候需要
}
