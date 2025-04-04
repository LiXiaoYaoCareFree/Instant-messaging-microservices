package user_models

import "IMM_server/common/models"

// FriendModel 好友表
type FriendModel struct {
	models.Model
	SendUserID    uint      `json:"sendUserID"`                     // 发起验证方
	SendUserModel UserModel `gorm:"foreignKey:SendUserID" json:"-"` // 发起验证方
	RevUserID     uint      `json:"revUserID"`                      // 接收验证方
	RevUserModel  UserModel `gorm:"foreignKey:RevUserID" json:"-"`  // 接收验证方
	Notice        string    `gorm:"size:128" json:"notice"`         // 备注                // 备注
}
