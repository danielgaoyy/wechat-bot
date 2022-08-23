package handlers

import (
	"errors"
	"fmt"
	"github.com/eatmoreapple/openwechat"
	"github.com/gaozicheng/workout/core"
	"regexp"
	"strconv"
	"strings"
)

var exerciseMembers []*member

const (
	checkIn   = `(打卡)\d+`
	showAll   = `查看进度`
	setTarget = `设目标\d+`
)

type member struct {
	*openwechat.User
	target  int64
	current int64
}

var group *openwechat.Group
var checkInRe, showRe, setRe, numberRe *regexp.Regexp

func InitExerciseGroup() {
	group = core.GetGroup("打卡")
	if group == nil {
		panic(errors.New("group not found"))
	}
	members, err := group.Members()
	if err != nil {
		panic(err)
	}
	for i := range members {
		exerciseMembers = append(exerciseMembers, &member{members[i], 0, 0})
	}
	checkInRe = regexp.MustCompile(checkIn)
	setRe = regexp.MustCompile(setTarget)
	showRe = regexp.MustCompile(showAll)
	numberRe = regexp.MustCompile(`\d+$`)
	fmt.Println(exerciseMembers)
}

func ProcessExercise(userName string, msg string) (string, error) {
	if match := checkInRe.MatchString(msg); match {
		progress := numberRe.FindString(msg)
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
		target := numberRe.FindString(msg)
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
	if match := showRe.MatchString(msg); match {
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
