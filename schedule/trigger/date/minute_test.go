package date

import "testing"

func TestMinute(t *testing.T) {
	m, err := New("10-20,30")
	if err != nil {
		t.Error(err)
		return
	}
	for i := 1; i < 10; i++ {
		if m.IsValid(i) {
			t.Error("分钟", i, "校验失败")
		}
	}
	for i := 10; i < 21; i++ {
		if !m.IsValid(i) {
			t.Error("分钟", i, "校验失败")
		}
	}
	for i := 21; i < 30; i++ {
		if m.IsValid(i) {
			t.Error("分钟", i, "校验失败")
		}
	}
	if !m.IsValid(30) {
		t.Error("分钟", 30, "校验失败")
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
