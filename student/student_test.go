package student

import (
	"log"
	"testing"
)

func TestStudent(t *testing.T) {
	if err := Load("./students.json"); err != nil {
		t.Error(err)
		return
	}
	defer func() {
		err := Save("./students.json")
		if err != nil {
			log.Fatalln(err)
		}
	}()
	if stu := GetStudentByQQ(1098105012); stu != nil {
		stu.AddTag("班委")
		t.Log(stu.ListTag())
		stu.RemoveTag("班委")
		t.Log(stu.ListTag())
		t.Log(stu.GetID(), stu.GetName())
	} else {
		t.Error("获取学生失败")
	}
}
