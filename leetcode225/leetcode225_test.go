package leetcode225

import "testing"

func TestMyStack1(t *testing.T) {
	obj := Constructor()
	obj.Push(1)
	obj.Push(2)
	if obj.top != 2 {
		t.Error("top failed")
	}
	pop := obj.Pop()
	if pop != 2 {
		t.Error("pop failed")
	}
	if obj.Empty() != false {
		t.Error("empty failed")
	}
}
func TestMyStack2(t *testing.T) {
	obj := Constructor()
	obj.Push(1)
	pop := obj.Pop()
	if pop != 1 {
		t.Error("pop failed")
	}
	if obj.Empty() != true {
		t.Error("empty failed")
	}
}
