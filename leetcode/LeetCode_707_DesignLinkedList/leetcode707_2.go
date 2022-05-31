package LeetCode_707_DesignLinkedList

//
//type listNode struct {
//	Val  int
//	Next *listNode
//}
//
//type MyLinkedList struct {
//	dummyhead *listNode
//	size      int
//}
//
//func Constructor() MyLinkedList {
//	return MyLinkedList{
//		dummyhead: &listNode{},
//		size:      0,
//	}
//}
//
//func (this *MyLinkedList) Get(index int) int {
//	if index > this.size-1 {
//		return -1
//	}
//	node := this.dummyhead
//	for i := 0; i <= index; i++ {
//		node = node.Next
//	}
//	return node.Val
//}
//
//func (this *MyLinkedList) AddAtHead(val int) {
//	node := &listNode{
//		Val:  val,
//		Next: nil,
//	}
//	node.Next = this.dummyhead.Next
//	this.dummyhead.Next = node
//	this.size++
//}
//
//func (this *MyLinkedList) AddAtTail(val int) {
//	cur := this.dummyhead
//	for i := 0; i < this.size; i++ {
//		cur = cur.Next
//	}
//	node := &listNode{
//		Val:  val,
//		Next: nil,
//	}
//	cur.Next = node
//	this.size++
//}
//
//func (this *MyLinkedList) AddAtIndex(index int, val int) {
//	if index > this.size {
//		return
//	}
//	prev, cur := this.dummyhead, this.dummyhead.Next
//	for i := 1; i <= index; i++ {
//		prev = cur
//		cur = cur.Next
//	}
//
//	node := &listNode{
//		Val:  val,
//		Next: nil,
//	}
//	prev.Next = node
//	node.Next = cur
//}
//
//func (this *MyLinkedList) DeleteAtIndex(index int) {
//	if index > this.size-1 {
//		return
//	}
//	prev, cur := this.dummyhead, this.dummyhead.Next
//	for i := 1; i <= index; i++ {
//		prev = cur
//		cur = cur.Next
//	}
//
//	prev.Next = cur.Next
//}
//
///**
// * Your MyLinkedList object will be instantiated and called as such:
// * obj := Constructor();
// * param_1 := obj.Get(index);
// * obj.AddAtHead(val);
// * obj.AddAtTail(val);
// * obj.AddAtIndex(index,val);
// * obj.DeleteAtIndex(index);
// */
