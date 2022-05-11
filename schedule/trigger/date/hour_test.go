package date

import "testing"

func TestHour(t *testing.T) {
	m, err := NewMinute("3-8,12")
	if err != nil {
		t.Error(err)
		return
	}
	for i := 0; i < 3; i++ {
		if m.IsValid(i) {
			t.Error("小时", i, "校验失败")
		}
	}
	for i := 3; i < 9; i++ {
		if !m.IsValid(i) {
			t.Error("小时", i, "校验失败")
		}
	}
	for i := 9; i < 12; i++ {
		if m.IsValid(i) {
			t.Error("小时", i, "校验失败")
		}
	}
	if !m.IsValid(12) {
		t.Error("小时", 30, "校验失败")
	}
	if m.GetNextMinute(12) != 13 {
		t.Error("获取12的下个时间点错误")
	}
	if m.GetNextMinute(25) != 30 {
		t.Error("获取25的下个时间点错误")
	}
	if m.GetNextMinute(-1) != 10 {
		t.Error("获取首个时间点错误")
	}
	if m.GetNextMinute(40) != -1 {
		t.Error("获取不存在的时间点错误")
	}
}
