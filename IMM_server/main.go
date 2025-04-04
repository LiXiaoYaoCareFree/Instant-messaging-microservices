package main

import (
	"IMM_server/core"
	"IMM_server/imm_chat/chat_models"
	"IMM_server/imm_group/group_models"
	"IMM_server/imm_user/user_models"
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
		db := core.InitGorm("")
		err := db.AutoMigrate(
			&user_models.UserModel{},         // 用户表
			&user_models.FriendModel{},       // 好友表
			&user_models.FriendVerifyModel{}, // 好友验证表
			&user_models.UserConfModel{},     // 用户配置表
			&chat_models.ChatModel{},         // 对话表
			&group_models.GroupModel{},       // 群组表
			&group_models.GroupMemberModel{}, // 群成员表
			&group_models.GroupMsgModel{},    // 群消息表
			&group_models.GroupVerifyModel{}, // 群验证表

		)
		if err != nil {
			fmt.Println("表结构生成失败", err)
			return
		}
		fmt.Println("表结构生成成功！")

	}

}
