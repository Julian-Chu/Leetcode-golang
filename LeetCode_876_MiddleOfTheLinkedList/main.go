package main

func main() {

}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
	var nodes []*ListNode
	node := head
	for {
		nodes = append(nodes, node)
		if node.Next == nil {
			break
		}
		node = node.Next
	}

	index := int(len(nodes) / 2)
	return nodes[index]
}

type ListNode struct {
	Val  int
	Next *ListNode
}
