package minute

import "testing"

func TestMinute(t *testing.T) {
	m, err := NewMinute("10-20,30")
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
}
