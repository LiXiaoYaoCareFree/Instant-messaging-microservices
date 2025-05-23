package user_models

import "IMM_server/common/models"

// UserModel 用户表
type UserModel struct {
	models.Model
	Pwd            string         `gorm:"size:64" json:"pwd"`
	Nickname       string         `gorm:"size:32" json:"nickname"`
	Abstract       string         `gorm:"size:128" json:"abstract"`
	Avatar         string         `gorm:"size:256" json:"avatar"`
	IP             string         `gorm:"size:32" json:"ip"`
	Addr           string         `gorm:"size:64" json:"addr"`
	Role           int8           `json:"role"`                          // 角色 1 管理员  2 普通用户
	OpenID         string         `gorm:"size:64" json:"token"`          // 第三方平台登录的凭证
	RegisterSource string         `gorm:"size:16" json:"registerSource"` // 注册来源
	UserConfModel  *UserConfModel `gorm:"foreignKey:UserID" json:"UserConfModel"`
}
