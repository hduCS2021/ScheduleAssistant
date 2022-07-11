package main

import (
	"context"
	"fmt"
	"github.com/BaiMeow/SimpleBot/message"
	"github.com/hduCS2021/ScheduleAssistant/qqbot"
	"github.com/hduCS2021/ScheduleAssistant/schedule"
	"github.com/hduCS2021/ScheduleAssistant/schedule/date"
	"github.com/hduCS2021/ScheduleAssistant/schedule/trigger"
	"github.com/hduCS2021/ScheduleAssistant/student"
	log "github.com/sirupsen/logrus"
	"time"
)

const dataPath = "./students.json"
const botPath = "ws://127.0.0.1:6700"
const botToken = ""

func Init() error {
	log.Info("Init.")
	if err := qqbot.Init(botPath, botToken); err != nil {
		return err
	}
	if err := student.Load(dataPath); err != nil {
		return err
	}
	date.SetBeginDate(time.Date(2022, time.August, 28, 0, 0, 0, 0, time.Local))
	log.Info("Init succeed.")
	return nil
}

func main() {
	if err := Init(); err != nil {
		log.Fatalf("Init fail：%v", err)
	}
	defer func(path string) {
		err := student.Save(path)
		if err != nil {
			log.Fatalf("Saving date fail: %v", err)
		}
	}(dataPath)
	ctx, cancel := context.WithCancel(context.Background())
	set := student.EmptySet().Add(student.GetStudentByQQ(1098105012))
	defer cancel()
	schedule.New("青年大学习", set,
		func(schedule *schedule.Schedule) {
			schedule.StuSet.Broadcast(message.New().Text("青年大学习提醒"))
		}).AddTrigger(trigger.New().Add(date.NewMinute("0-59"))).Run(ctx)
	schedule.New("青年大学习结算", set,
		func(schedule *schedule.Schedule) {
			//export
			var list []string
			schedule.StuSet.ForEach(func(stu student.Student) bool {
				list = append(list, stu.GetName())
				return true
			})
			student.GetStudentByQQ(1098105012).Send(message.New().Text(fmt.Sprint(list)))
			//reset
			schedule.StuSet.Add(student.GetStudentByQQ(1098105012))
			//student.All().ForEach(func(stu student.Student) bool {
			//	schedule.StuSet.Add(stu)
			//	return true
			//})
		}).AddTrigger(trigger.New().Add(date.Weekday(7), date.Hour(22), date.Minute(0))).Run(ctx)
	qqbot.AddMsgHandler(func(stu student.Student, msg message.Msg) bool {
		if len(msg) != 1 || msg[0].GetType() != "image" {
			return false
		}
		_ = msg[0].(message.Image)
		if true {
			set.Remove(stu)
			stu.Send(message.New().Text("青年大学习打卡成功"))
		}
		return true
	})
	select {}
}
