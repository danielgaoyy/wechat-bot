package core

import (
	"fmt"
	"github.com/eatmoreapple/openwechat"
)

var Bot *openwechat.Bot
var Self *openwechat.Self
var Friends openwechat.Friends
var Groups openwechat.Groups

func Init() {
	var err error
	Bot = openwechat.DefaultBot()
	if err = Bot.Login(); err != nil {
		fmt.Println(err)
		panic(err)
	}
	Self, err = Bot.GetCurrentUser()
	if err != nil {
		panic(err)
	}
	Friends, err = Self.Friends()
	if err != nil {
		panic(err)
	}
	Groups, err = Self.Groups()
	if err != nil {
		panic(err)
	}
	fmt.Println(Groups)
}

func GetUser(userName string) *openwechat.Friend {
	return Friends.GetByUsername(userName)
}

func GetGroup(remarkName string) *openwechat.Group {
	return Groups.GetByRemarkName(remarkName)
}
