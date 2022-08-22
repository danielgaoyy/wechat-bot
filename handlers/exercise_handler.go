package handlers

import (
	"code.byted.org/gaozicheng/workout/core"
	"errors"
	"fmt"
	"github.com/eatmoreapple/openwechat"
	"regexp"
	"strconv"
	"strings"
)

var exerciseMembers []*member

const (
	checkIn   = `(打卡)\d+`
	showAll   = "查看进度"
	setTarget = `设目标\d+`
)

type member struct {
	*openwechat.User
	target  int64
	current int64
}

var checkInRe, setRe *regexp.Regexp

func InitExerciseGroup() {
	members, err := core.GetGroup("打卡").Members()
	if err != nil {
		panic(err)
	}
	for i := range members {
		exerciseMembers = append(exerciseMembers, &member{members[i], 0, 0})
	}
	checkInRe = regexp.MustCompile(checkIn)
	setRe = regexp.MustCompile(setTarget)
	fmt.Println(exerciseMembers)
}

func ProcessExercise(userName string, msg string) (string, error) {
	if match := checkInRe.MatchString(msg); match {
		progress := checkInRe.FindString(`^\d+$`)
		realProgress, err := strconv.ParseInt(progress, 10, 64)
		if err != nil {
			return "", err
		}
		err = UpdateCurrent(userName, realProgress)
		if err != nil {
			return "", err
		}
		return "更新成功", err
	}

	if match := setRe.MatchString(msg); match {
		target := setRe.FindString(`^\d+$`)
		realTarget, err := strconv.ParseInt(target, 10, 64)
		if err != nil {
			return "", err
		}
		err = SetTarget(userName, realTarget)
		if err != nil {
			return "", err
		}
		return "目标设置成功", err
	}
	if msg == showAll {
		return GetCurrent(), nil
	}
	return "", errors.New("unrecognized")
}

func SetTarget(userName string, target int64) error {
	for i := range exerciseMembers {
		if exerciseMembers[i].UserName == userName {
			exerciseMembers[i].target = target
			return nil
		}
	}
	return errors.New("user not found")
}

func UpdateCurrent(userName string, progress int64) error {
	for i := range exerciseMembers {
		if exerciseMembers[i].UserName == userName {
			exerciseMembers[i].current += progress
			return nil
		}
	}
	return errors.New("user not found")
}

func GetCurrent() string {
	ret := make([]string, 0)
	for _, mem := range exerciseMembers {
		ret = append(ret, fmt.Sprintf("%v:\n本周目标:%v, 当前进度:%v", mem.NickName, mem.target, mem.current))
	}
	return strings.Join(ret, "\n")
}
