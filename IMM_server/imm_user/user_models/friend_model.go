package user_models

import (
	"IMM_server/common/models"
	"gorm.io/gorm"
)

// FriendModel 好友表
type FriendModel struct {
	models.Model
	SendUserID    uint      `json:"sendUserID"`                     // 发起验证方
	SendUserModel UserModel `gorm:"foreignKey:SendUserID" json:"-"` // 发起验证方
	RevUserID     uint      `json:"revUserID"`                      // 接收验证方
	RevUserModel  UserModel `gorm:"foreignKey:RevUserID" json:"-"`  // 接收验证方
	SenUserNotice string    `gorm:"size:128" json:"senUserNotice"`  // 发送方备注
	RevUserNotice string    `gorm:"size:128" json:"revUserNotice"`  // 接收方备注
}

/*
A -> B  a发起添加b的好友请求

SenUserNotice就是 A对B的好友备注
RevUserNotice就是 B对A的好友备注
*/

func (f *FriendModel) IsFriend(db *gorm.DB, A, B uint) bool {
	err := db.Take(&f, "(send_user_id = ? and rev_user_id = ? ) or (send_user_id = ? and rev_user_id = ? )", A, B, B, A).Error
	if err == nil {
		return true
	}
	return false
}

func (f *FriendModel) GetUserNotice(userID uint) string {
	if userID == f.SendUserID {
		// 如果我是发起方
		return f.SenUserNotice
	}
	if userID == f.RevUserID {
		// 如果我是接收方
		return f.RevUserNotice
	}
	return ""
}
