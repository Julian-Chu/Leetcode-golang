package solution2

import "testing"

func TestCase1(t *testing.T) {
	queue := Constructor()
	queue.Push(1)
	queue.Push(2)
	if res := queue.Peek(); res != 1 {
		t.Errorf("peek failed, should be %d, but get %d", 1, res)
	}
	if res := queue.Pop(); res != 1 {
		t.Errorf("peek failed, should be %d, but get %d", 1, res)
	}
	if res := queue.Empty(); res != false {
		t.Errorf("peek failed, should be %t, but get %t", false, res)
	}
}
