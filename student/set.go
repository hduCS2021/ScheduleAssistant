package student

import (
	"github.com/BaiMeow/SimpleBot/message"
	"sync"
)

type StuSet struct {
	students []Student
	lock     *sync.RWMutex
}

//EmptySet returns a new empty StuSet
func EmptySet() *StuSet {
	return &StuSet{
		students: nil,
		lock:     new(sync.RWMutex),
	}
}

//All returns a StuSet that contains all students
func All() *StuSet {
	set := EmptySet()
	for index := range allStudents {
		set.students = append(set.students, &allStudents[index])
	}
	return set
}

//FilterTag returns a new StuSet including students that are tagged by certain tag and are in original StuSet
func (set *StuSet) FilterTag(tag string) *StuSet {
	set.lock.RLock()
	defer set.lock.RUnlock()
	set1 := EmptySet()
	for _, v := range set.students {
		if v.ExistTag(tag) {
			set1.students = append(set1.students, v)
		}
	}
	return set1
}

//FilterFunc returns a new StuSet including Students that satisfied func f and are in original StuSet
func (set *StuSet) FilterFunc(f func(Student) bool) *StuSet {
	set.lock.Lock()
	defer set.lock.Unlock()
	set1 := EmptySet()
	for _, v := range set.students {
		if f(v) {
			set1.students = append(set1.students, v)
		}
	}
	return set1
}

//Add a Student to the Set if not exists, and return the original StuSet
func (set *StuSet) Add(stu Student) *StuSet {
	set.lock.Lock()
	defer set.lock.Unlock()
	if !set.existWithoutLock(stu) {
		set.students = append(set.students, stu)
	}
	return set
}

//Remove a Student to the Set if exists, and return the original StuSet
func (set *StuSet) Remove(stu Student) {
	set.lock.Lock()
	defer set.lock.Unlock()
	i := 0
	for i = 0; i < len(set.students); i++ {
		if (set.students)[i].GetQQ() == stu.GetQQ() {
			if i == len(set.students)-1 {
				set.students = (set.students)[:i]
			} else {
				set.students = append((set.students)[:i], (set.students)[i+1:]...)
			}
			return
		}
	}
}

//Exist check if a Student is included in a StuSet
func (set *StuSet) Exist(stu Student) bool {
	set.lock.RLock()
	defer set.lock.RUnlock()
	return set.existWithoutLock(stu)
}

func (set *StuSet) existWithoutLock(stu Student) bool {
	for _, v := range set.students {
		if v.GetQQ() == stu.GetQQ() {
			return true
		}
	}
	return false
}

//Broadcast broadcasts a message to all Students in StuSet
func (set *StuSet) Broadcast(msg message.Msg) {
	set.ForEach(func(stu Student) bool {
		stu.Send(msg)
		return false
	})
}

//ForEach do func to every student in StuSet,
//break when F returns false
func (set *StuSet) ForEach(F func(stu Student) bool) {
	set.lock.Lock()
	defer set.lock.Unlock()
	for i := range set.students {
		if !F(set.students[i]) {
			break
		}
	}
}
