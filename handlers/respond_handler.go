package handlers

import (
	"code.byted.org/gaozicheng/workout/core"
	"fmt"
	"github.com/eatmoreapple/openwechat"
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
		if !msg.IsSendBySelf() {
			sender, err = msg.SenderInGroup()
		} else {
			sender = core.Self.User
		}

		if err != nil {
			fmt.Println("get sender failed")
			return
		}
		ret, err := ProcessExercise(sender.UserName, msg.RowContent)
		if err != nil {
			fmt.Printf("process exercise failed:%v\n", err)
			return
		}
		_, _ = msg.ReplyText(ret)
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
