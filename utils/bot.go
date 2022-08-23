package utils

import (
	"fmt"
	"github.com/kevwan/chatbot/bot"
	"github.com/kevwan/chatbot/bot/adapters/logic"
	"github.com/kevwan/chatbot/bot/adapters/storage"
	"log"
	"path/filepath"
	"strings"
)

var (
	dir       = "../data"
	storeFile = "../data/corpus.gob"
	tops      = 1
	chatBot   *bot.ChatBot
)

func init() {

	files := findCorporaFiles(dir)

	var corporaFiles string

	corporaFiles = strings.Join(files, ",")

	if len(corporaFiles) == 0 {
		return
	}
	store, err := storage.NewSeparatedMemoryStorage(storeFile)
	if err != nil {
		panic(err)
	}
	chatBot = &bot.ChatBot{
		Trainer:      bot.NewCorpusTrainer(store),
		LogicAdapter: logic.NewClosestMatch(store, tops),
	}
	if err := chatBot.Train(strings.Split(corporaFiles, ",")); err != nil {
		log.Fatal(err)
	}
}

func GetResponse(msg string) string {
	resp := chatBot.GetResponse(msg)
	if len(resp) > 0 {
		return resp[0].Content
	}
	return ""
}

func findCorporaFiles(dir string) []string {
	var files []string

	jsonFiles, err := filepath.Glob(filepath.Join(dir, "*.json"))
	if err != nil {
		fmt.Println(err)
		return nil
	}

	files = append(files, jsonFiles...)

	ymlFiles, err := filepath.Glob(filepath.Join(dir, "*.yml"))
	if err != nil {
		fmt.Println(err)
		return nil
	}

	files = append(files, ymlFiles...)

	yamlFiles, err := filepath.Glob(filepath.Join(dir, "*.yaml"))
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return append(files, yamlFiles...)
}
