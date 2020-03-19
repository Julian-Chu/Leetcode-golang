package leetcode385

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
	b := []byte(s)
	ni := &NestedInteger{}
	if b[0] != '[' {
		val := 0
		for i := range b {
			val = val*10 + int(b[i]-'0')
		}
		ni.SetInteger(val)
		return ni
	}

	// remove first  [ and ]
	b = b[1 : len(b)-1]
	n := len(b)
	i := 0
	j := 0
	cnt := 0

	for j < n {
		switch {
		case b[j] == '[':
			cnt++
			j++
		case b[j] >= '0' && b[j] <= '9':
			j++
			if j == n {
				ni.Add(*deserialize(string(b[i:j])))
			}
		case b[j] == ',' && cnt == 0:
			ni.Add(*deserialize(string(b[i:j])))
			j++
			i = j
		case b[j] == ']':
			j++
			cnt--
			if cnt == 0 {
				ni.Add(*deserialize(string(b[i:j])))
			}
		}
	}

	return ni
}
