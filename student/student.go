package student

import (
	"github.com/BaiMeow/SimpleBot/message"
	"github.com/hduCS2021/ScheduleAssistant/qqbot"
	"log"
	"sync"
)

type Student interface {
	GetName() string
	GetQQ() int64
	GetID() int32
	ExistTag(tag string) bool
	AddTag(tag string)
	RemoveTag(tag string)
	ListTag() []string
	Send(msg message.Msg)
}

type student struct {
	ID   int32    `json:"id"`
	Name string   `json:"name"`
	QQ   int64    `json:"qq"`
	Tags []string `json:"tags"`
	lock sync.RWMutex
}

var allStudents []student

func (stu *student) GetName() string {
	return stu.Name
}

func (stu *student) GetQQ() int64 {
	return stu.QQ
}

func (stu *student) GetID() int32 {
	return stu.ID
}

// ExistTag return true when tag exists.
func (stu *student) ExistTag(tag string) bool {
	stu.lock.RLock()
	defer stu.lock.RUnlock()
	return stu.existTagWithoutLock(tag)
}

func (stu *student) existTagWithoutLock(tag string) bool {
	for _, v := range stu.Tags {
		if v == tag {
			return true
		}
	}
	return false
}

func (stu *student) Send(msg message.Msg) {
	for i := 0; i < 3; i++ {
		err := qqbot.SendMessage(stu.QQ, msg)
		if err != nil {
			log.Printf("发送消息失败（qq=%d）:%v\n", stu.QQ, err)
			continue
		}
		return
	}
}

func (stu *student) AddTag(tag string) {
	if !stu.existTagWithoutLock(tag) {
		stu.Tags = append(stu.Tags, tag)
	}
}

func (stu *student) ListTag() []string {
	stu.lock.RLock()
	defer stu.lock.RUnlock()
	var tmp = make([]string, len(stu.Tags))
	copy(tmp, stu.Tags)
	return tmp
}

func (stu *student) RemoveTag(tag string) {
	stu.lock.Lock()
	defer stu.lock.Unlock()
	for i, v := range stu.Tags {
		if v == tag {
			if i == len(stu.Tags)-1 {
				stu.Tags = stu.Tags[:i]
			} else {
				stu.Tags = append(stu.Tags[:i], stu.Tags[i+1:]...)
			}
			return
		}
	}

}
