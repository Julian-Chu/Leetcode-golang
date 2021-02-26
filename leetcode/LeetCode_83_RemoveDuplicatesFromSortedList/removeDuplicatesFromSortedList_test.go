package LeetCode_83_RemoveDuplicatesFromSortedList

import (
	"testing"
)

func Test_LinkedListEqual1(t *testing.T) {
	head := &ListNode{1, nil}
	expected := &ListNode{1, nil}
	if !LinkedListEqual(deleteDuplicates(head), expected) {
		t.Error("Not match")
	}
}
func Test_LinkedListEqual2(t *testing.T) {
	next := &ListNode{2, nil}
	head := &ListNode{1, next}
	expected := &ListNode{1, next}
	if !LinkedListEqual(deleteDuplicates(head), expected) {
		t.Error("Not match")
	}
}

func LinkedListEqual(actual, expected *ListNode) bool {

	for actual.Val == expected.Val {
		if actual.Next == nil && expected.Next == nil {
			return true
		} else if actual.Next == nil || expected.Next == nil {
			break
		}
		actual = actual.Next
		expected = expected.Next

	}
	return false
}

func Test_Case1(t *testing.T) {
	secondNode := &ListNode{2, nil}
	firstNode := &ListNode{1, secondNode}
	head := &ListNode{1, firstNode}

	expected := &ListNode{1, secondNode}
	if !LinkedListEqual(deleteDuplicates(head), expected) {
		t.Error("Not Match")
	}
}
func Test_Case2(t *testing.T) {
	forthNode := &ListNode{3, nil}
	thirdNode := &ListNode{3, forthNode}
	secondNode := &ListNode{2, thirdNode}
	firstNode := &ListNode{1, secondNode}
	head := &ListNode{1, firstNode}

	expected := &ListNode{1, &ListNode{2, &ListNode{3, nil}}}
	if !LinkedListEqual(deleteDuplicates(head), expected) {
		t.Error("Not Match")
	}
}
