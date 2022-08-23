package main

import (
	"github.com/gaozicheng/workout/core"
	"github.com/gaozicheng/workout/handlers"
)

func main() {
	core.Init()
	handlers.InitExerciseGroup()
	handlers.InitUniqueHandler()
	handlers.InitRespondHandler()
	_ = core.Bot.Block()
}
