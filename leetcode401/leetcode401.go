package leetcode401

import (
	"fmt"
	"math/bits"
)

func readBinaryWatch(num int) []string {
	res := make([]string, 0, 10)

	for h := uint(0); h < 12; h++ {
		for m := uint(0); m < 60; m++ {
			if bits.OnesCount(h<<6+m) == num {
				res = append(res, fmt.Sprintf("%d:%02d", h, m))
			}
		}
	}
	return res
}
