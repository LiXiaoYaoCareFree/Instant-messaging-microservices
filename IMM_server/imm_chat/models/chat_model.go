package models

import (
	"IMM_server/common/models"
	"time"
)

type ChatModel struct {
	models.Model
	SendUserID uint       `json:"sendUserID"`
	RevUserID  uint       `json:"revUserID"`
	MsgType    int8       `json:"msgType"`    // 消息类型 1 文本类型 2 图片消息 3 视频消息 4 文件消息 5 语音消息 6 语音通话 7 视频通话 8 撤回消息 9 回复消息 10 引用消息
	MsgPreview string     `json:"msgPreview"` // 消息预览
	Msg        Msg        `json:"msg"`        // 消息内容
	SystemMsg  *SystemMsg `json:"systemMsg"`  // 系统提示
}

type SystemMsg struct {
	Type int8 `json:"type"` // 违规类型 1 涉黄 2 涉恐 3 涉证 4 不正当言论
}

type Msg struct {
	Type         int8          `json:"type"`         // 消息类型 和 msgType一模一样
	Content      *string       `json:"content"`      // 为1时使用
	ImageMsg     *ImageMsg     `json:"imageMsg"`     // 图片消息
	VideoMsg     *VideoMsg     `json:"videoMsg"`     // 视频消息
	FileMsg      *FileMsg      `json:"fileMsg"`      // 文件消息
	VoiceMsg     *VoiceMsg     `json:"voiceMsg"`     // 语音消息
	VoiceCallMsg *VoiceCallMsg `json:"voiceCallMsg"` // 语音通话
	VideoCallMsg *VideoCallMsg `json:"videoCallMsg"` // 视频通话
	WithdrawMsg  *WithdrawMsg  `json:"withdrawMsg"`  // 撤回消息
	ReplyMsg     *ReplyMsg     `json:"replyMsg"`     // 9回复消息
	QuoteMsg     *QuoteMsg     `json:"quoteMsg"`     // 引用消息
}

type ImageMsg struct {
	Title string `json:"title"`
	Src   string `json:"src"`
}

type VideoMsg struct {
	Title string `json:"title"`
	Src   string `json:"src"`
	Time  int    `json:"time"` // 时长 单位秒
}

type FileMsg struct {
	Title string `json:"title"`
	Src   string `json:"src"`
	Size  int64  `json:"size"` // 文件大小
	Type  string `json:"type"` // 文件类型 word
}

type VoiceMsg struct {
	Src  string `json:"src"`
	Time int    `json:"time"` // 时长 单位秒
}

type VoiceCallMsg struct {
	StartTime time.Time `json:"startTime"` // 开始时间
	EndTime   time.Time `json:"endTime"`   // 结束时间
	EndReason int8      `json:"endReason"` // 结束原因 0 发起方挂断 1 接收方挂断 2 网络原因挂断 3 未打通
}

type VideoCallMsg struct {
	StartTime time.Time `json:"startTime"` // 开始时间
	EndTime   time.Time `json:"endTime"`   // 结束时间
	EndReason int8      `json:"endReason"` // 结束原因 0 发起方挂断 1 接收方挂断 2 网络原因挂断 3 未打通
}

// WithdrawMsg 撤回消息
type WithdrawMsg struct {
	Content   string `json:"content"` // 撤回的提示词
	OriginMsg *Msg   `json:"-"`       // 原消息
}
type ReplyMsg struct {
	MsgID   uint   `json:"msgID"`   // 消息id
	Content string `json:"content"` // 回复的文本消息，目前只能限制回复文本
	Msg     *Msg   `json:"msg"`
}

type QuoteMsg struct {
	MsgID   uint   `json:"msgID"`   // 消息id
	Content string `json:"content"` // 回复的文本消息，目前只能限制回复文本
	Msg     *Msg   `json:"msg"`
}
