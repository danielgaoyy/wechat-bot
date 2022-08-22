package handlers

import "fmt"

const (
	startRepeat = "开启自动回复"
	endRepeat   = "关闭自动回复"
)

type uniqueHandler struct {
	repeat bool
}

var unique *uniqueHandler

func InitUniqueHandler() {
	unique = &uniqueHandler{repeat: false}
}

func ProcessUnique(msg string) (string, bool) {
	ret, respond := SwitchRepeat(msg)
	if respond {
		return ret, respond
	}
	return Repeat(msg)
}

func SwitchRepeat(msg string) (string, bool) {
	var retS string
	var toRet bool
	switch msg {
	case startRepeat:
		unique.repeat = true
		retS = "自动回复已开启"
		toRet = true
	case endRepeat:
		unique.repeat = true
		retS = "自动回复已关闭"
		toRet = true
	}
	return retS, toRet
}

func Repeat(msg string) (string, bool) {
	if unique.repeat {
		return fmt.Sprintf("收到你的消息：%v\n暂时无法回复，请见谅", msg), true
	}
	return "", false
}
