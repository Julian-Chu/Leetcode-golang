package main

import (
	"testing"
)

func TestMiddleNode_OneNde(t *testing.T) {
	node := ListNode{
		1, nil,
	}
	adr := middleNode(&node)

	if (*adr).Val != 1 {
		t.Error("not pass")
	}
}

func TestMiddleNode_TwoNodes(t *testing.T) {
	node2 := ListNode{
		2, nil,
	}
	node1 := ListNode{
		1, &node2,
	}

	adr := middleNode(&node1)

	if (*adr).Val != 2 {
		t.Error("not pass")
	}
}
