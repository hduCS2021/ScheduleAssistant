package student

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func Load(path string) error {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return fmt.Errorf("无法打开学生数据文件:%v", err)
	}
	var students []student
	if err := json.Unmarshal(bytes, &students); err != nil {
		return fmt.Errorf("学生数据文件解析失败:%v", err)
	}
	allStudents = make([]student, len(students))
	copy(allStudents, students)
	return nil
}

func Save(path string) error {
	bytes, err := json.Marshal(&allStudents)
	if err != nil {
		return fmt.Errorf("序列化学生数据失败：%v", err)
	}
	if err := ioutil.WriteFile(path, bytes, 0777); err != nil {
		return fmt.Errorf("写入文件失败:%v", err)
	}
	return nil
}

func GetStudentByID(id int32) Student {
	return GetStudentFunc(func(s Student) bool {
		return s.GetID() == id
	})
}

func GetStudentByName(name string) Student {
	return GetStudentFunc(func(s Student) bool {
		return s.GetName() == name
	})
}

func GetStudentByQQ(qq int64) Student {
	return GetStudentFunc(func(s Student) bool {
		return s.GetQQ() == qq
	})
}

func GetStudentFunc(f func(Student) bool) Student {
	for _, v := range All().students {
		if f(v) {
			return v
		}
	}
	return nil
}
