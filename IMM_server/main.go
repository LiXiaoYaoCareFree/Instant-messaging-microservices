package main

import (
	"IMM_server/core"
	models2 "IMM_server/imm_chat/models"
	models3 "IMM_server/imm_group/models"
	"IMM_server/imm_user/models"
	"flag"
	"fmt"
)

type Options struct {
	DB bool
}

func main() {

	var opt Options
	flag.BoolVar(&opt.DB, "db", false, "db")
	flag.Parse()

	if opt.DB {
		db := core.InitMysql()
		err := db.AutoMigrate(
			&models.UserModel{},
			&models.FriendModel{},
			&models.FriendVerifyModel{},
			&models.UserConfModel{},

			&models2.ChatModel{},
			&models3.GroupModel{},
			&models3.GroupMemberModel{},
			&models3.GroupMsgModel{},
			&models3.GroupVerifyModel{},
		)
		if err != nil {
			fmt.Println("表结构生成失败", err)
			return
		}
		fmt.Println("表结构生成成功！")

	}

}
