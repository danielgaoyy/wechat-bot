package handlers

import (
	"fmt"
	"github.com/eatmoreapple/openwechat"
	"github.com/gaozicheng/workout/core"
)

func InitRespondHandler() {
	core.Bot.MessageHandler = MakeResponse
}

func MakeResponse(msg *openwechat.Message) {
	sender, err := msg.Sender()
	fmt.Println(sender.RemarkName)
	if err != nil {
		return
	}
	if msg.IsComeFromGroup() {
		if sender.NickName != "打卡" && !msg.IsSendBySelf() {
			return
		}
		fmt.Println("hit group chat")
		var userName string
		if !msg.IsSendBySelf() {
			sender, err = msg.SenderInGroup()
			userName = sender.UserName
		} else {
			userName = "高小毛"
		}

		if err != nil {
			fmt.Println("get sender failed")
			return
		}
		ret, err := ProcessExercise(userName, msg.RowContent)
		if err != nil {
			fmt.Printf("process exercise failed:%v\n", err)
			return
		}
		if msg.IsSendBySelf() {
			_, _ = group.SendText(ret)
		} else {
			_, _ = msg.ReplyText(ret)
		}
		return
	}
	if sender.RemarkName == "唐志卿" {
		fmt.Println("user name hit")
		ret, send := ProcessUnique(msg.RowContent)
		if send || ret != "" {
			_, _ = msg.ReplyText(ret)
		}
	}
}
