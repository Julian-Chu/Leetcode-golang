package leetcode397

func integerReplacement(n int) int {
	rec := make(map[int]int)
	rec[1] = 0

	var ir func(int) int
	ir = func(i int) int {
		if n, ok := rec[i]; ok {
			return n
		}

		if i%2 == 0 {
			rec[i] = ir(i/2) + 1
			return rec[i]
		}

		rec[i] = min(ir(i+1), ir(i-1)) + 1
		return rec[i]
	}

	return ir(n)
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
