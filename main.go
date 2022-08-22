package main

import (
	"fmt"
	"github.com/eatmoreapple/openwechat"
)

func main() {
	bot := openwechat.DefaultBot()
	if err := bot.Login(); err != nil {
		fmt.Println(err)
		return
	}
	err := bot.Block()
	if err != nil {
		fmt.Printf("bot error:%v\n", err)
	}
}
