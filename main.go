package main

import (
	"code.byted.org/gaozicheng/workout/core"
	"code.byted.org/gaozicheng/workout/handlers"
)

func main() {
	core.Init()
	handlers.InitExerciseGroup()
	handlers.InitUniqueHandler()
	handlers.InitRespondHandler()
	_ = core.Bot.Block()
}
