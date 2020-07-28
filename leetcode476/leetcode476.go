package leetcode476

func findComplement(num int) int {
	comp := ^num

	cnt := 0
	for num != 0 {
		cnt++
		num >>= 1
	}

	mask := 1<<cnt - 1

	return comp & mask
}
