package leetcode385

import (
	"strconv"
)

//// This is the interface that allows for creating nested lists.
//// You should not implement it, or speculate about its implementation
type NestedInteger struct {
	val  int
	list []*NestedInteger
}

//
//// Return true if this NestedInteger holds a single integer, rather than a nested list.
func (n NestedInteger) IsInteger() bool {
	return n.list == nil
}

//
//// Return the single integer that this NestedInteger holds, if it holds a single integer
//// The result is undefined if this NestedInteger holds a nested list
//// So before calling this method, you should have a check
func (n NestedInteger) GetInteger() int {
	return n.val
}

//
//// Set this NestedInteger to hold a single integer.
func (n *NestedInteger) SetInteger(value int) {
	n.val = value
}

//
//// Set this NestedInteger to hold a nested list and adds a nested integer to it.
func (n *NestedInteger) Add(elem NestedInteger) {
	n.list = append(n.list, &elem)
}

//
//// Return the nested list that this NestedInteger holds, if it holds a nested list
//// The list length is zero if this NestedInteger holds a single integer
//// You can access NestedInteger's List element directly if you want to modify it
func (n NestedInteger) GetList() []*NestedInteger {
	return n.list
}

func deserialize(s string) *NestedInteger {
	if len(s) == 0 {
		return nil
	}

	if s[0] != '[' {
		return getValue(s)
	}

	stack := new(stackChars)

	var cur *NestedInteger

	si, ci := 0, 0

	for ; ci < len(s); ci++ {
		switch s[ci] {
		case '[':
			if cur != nil {
				stack.Push(cur)
			}

			cur = new(NestedInteger)
			si = ci + 1
		case ']':
			if ci > si {
				cur.Add(*getValue(s[si:ci]))
			}

			if !stack.Empty() {
				tmp := stack.Pop()
				tmp.Add(*cur)
				cur = tmp
			}

			si = ci + 1
		case ',':
			if s[ci-1] != ']' {
				cur.Add(*getValue(s[si:ci]))
			}
			si = ci + 1
		}
	}

	return cur
}

func getValue(s string) *NestedInteger {
	val, _ := strconv.Atoi(s)
	item := new(NestedInteger)
	item.SetInteger(val)
	return item
}

type stackChars struct {
	chars []*NestedInteger
}

func (s *stackChars) Pop() *NestedInteger {
	slen := len(s.chars)
	rb := s.chars[slen-1]
	s.chars = s.chars[:slen-1]
	return rb
}

func (s *stackChars) Push(nb *NestedInteger) {
	s.chars = append(s.chars, nb)
}
func (s stackChars) Empty() bool {
	return len(s.chars) == 0
}
