package student

import (
	"testing"
)

func TestStuSet(t *testing.T) {
	if err := Load("./students.json"); err != nil {
		t.Error(err)
		return
	}
	defer func() {
		err := Save("./students.json")
		if err != nil {
			t.Fatal(err)
		}
	}()
	stus := All()
	stus.ForEach(func(stu *Student) bool {
		t.Log((*stu).GetName())
		return true
	})
	stu := GetStudentByQQ(1098105012)
	stu.AddTag("123")
	stus.FilterTag("123")
	if !stus.Exist(stu) {
		t.Error("filter error")
	}
	stu.RemoveTag("123")
	stus = All().FilterTag("123")
	stus.ForEach(func(stu *Student) bool {
		t.Log((*stu).GetName())
		return true
	})
}
