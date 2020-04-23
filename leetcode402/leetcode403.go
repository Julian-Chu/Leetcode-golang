package leetcode402

import (
	"strconv"
	"strings"
)

func removeKdigits(num string, k int) string {
	if len(num) == 0 {
		return "0"
	}
	if k == 0 {
		return num
	}

	bs := []byte(num)
	for j := 0; j < k; j++ {
		if len(bs) == 0 {
			return "0"
		}
		zeros := 0
		for i := 1; i < len(bs); i++ {
			if bs[i] != '0' {
				break
			}
			zeros++
		}

		if zeros > 0 {
			bs = bs[zeros+1:]
			continue
		}

		min := bs[1:]
		for i := 1; i < len(bs); i++ {
			target := make([]byte, len(bs)-1)
			copy(target, bs[:i])
			copy(target[i:], bs[i+1:])

			minNum, _ := strconv.Atoi(string(min))
			targetNum, _ := strconv.Atoi(string(target))
			if minNum > targetNum {
				min = target
			}
		}
		bs = min
	}
	v, _ := strconv.Atoi(string(bs))
	if len(bs) == 0 || v == 0 {
		return "0"
	}
	return strings.TrimLeft(string(bs), "0")
}
