package qqbot

import (
	"github.com/BaiMeow/SimpleBot/bot"
	"github.com/BaiMeow/SimpleBot/driver"
	"github.com/BaiMeow/SimpleBot/message"
	"github.com/hduCS2021/ScheduleAssistant/student"
)

var b *bot.Bot
var handler []func(stu student.Student, msg message.Msg) bool

func Init(addr, token string) error {
	b = bot.New(driver.NewWsDriver(addr, token))
	b.Attach(&bot.PrivateMsgHandler{
		F: func(MsgID int32, UserID int64, Msg message.Msg) bool {
			stu := student.GetStudentByQQ(UserID)
			if stu == nil {
				return false
			}
			for index := range handler {
				if handler[index](stu, Msg) {
					break
				}
			}
			return true
		}})
	b.Attach(&bot.FriendAddHandler{
		F: func(request *bot.FriendRequest) bool {
			if stu := student.GetStudentByQQ(request.UserID); stu != nil {
				request.Agree(stu.GetName())
				return true
			}
			return false
		}})
	return b.Run()
}

func SendMessage(qq int64, msg message.Msg) error {
	_, err := b.SendPrivateMsg(qq, msg)
	return err
}

func AddMsgHandler(F func(stu student.Student, msg message.Msg) bool) {
	handler = append(handler, F)
}
