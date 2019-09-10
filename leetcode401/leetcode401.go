package leetcode401

import "fmt"

func readBinaryWatch(num int) []string {
	res := make([]string, 0, 10)
	bits := make([]bool, 10)
	var dfs func(int, int)
	dfs = func(index int, num int) {
		if num == 0 {
			m, h := get(bits[:6]), get(bits[6:])
			if h < 12 && m < 60 {
				res = append(res, fmt.Sprintf("%d:%02d", h, m))
			}
			return
		}

		for i := index; i < 10-num+1; i++ {
			bits[i] = true
			dfs(i+1, num-1)
			bits[i] = false
		}

	}
	dfs(0, num)
	return res
}

var values = []int{1, 2, 4, 8, 16, 32}

func get(bits []bool) int {
	sum := 0
	for k, v := range bits {
		if v == true {
			sum += values[k]
		}
	}
	return sum
}

//func readBinaryWatch(num int) []string {
//	res := make([]string, 0, 10)
//
//	for h := uint(0); h < 12; h++ {
//		for m := uint(0); m < 60; m++ {
//			if bits.OnesCount(h<<6+m) == num {
//				res = append(res, fmt.Sprintf("%d:%02d", h, m))
//			}
//		}
//	}
//	return res
//}
//
