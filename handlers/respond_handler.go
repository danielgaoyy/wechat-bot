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
	if msg.IsComeFromGroup() && msg.IsAt() && sender.RemarkName == "健身打卡" {
		sender, err = msg.SenderInGroup()
		if err != nil {
			return
		}
		ret, err := ProcessExercise(sender.UserName, msg.RowContent)
		if err != nil {
			return
		}
		_, _ = msg.ReplyText(ret)
		return
	}
	if sender.RemarkName == "唐志卿" || sender.RemarkName == "孙子" {
		fmt.Println("user name hit")
		ret, send := ProcessUnique(msg.RowContent)
		if send || ret != "" {
			_, _ = msg.ReplyText(ret)
		}
	}
}
