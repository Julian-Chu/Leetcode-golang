package leetcode405

import "strings"

func toHex(num int) string {
	if num == 0 {
		return "0"
	}
	m := map[int]byte{
		0:  '0',
		1:  '1',
		2:  '2',
		3:  '3',
		4:  '4',
		5:  '5',
		6:  '6',
		7:  '7',
		8:  '8',
		9:  '9',
		10: 'a',
		11: 'b',
		12: 'c',
		13: 'd',
		14: 'e',
		15: 'f',
	}

	res := make([]byte, 8)
	for i := 7; i >= 0; i-- {
		v, _ := m[num&15]
		res[i] = v
		num = num >> 4
	}

	return strings.TrimLeft(string(res), "0")
}
