package leetcode402

func removeKdigits(num string, k int) string {
	if len(num) <= k {
		return "0"
	}

	stack := make([]byte, len(num))
	top := 0
	digits := len(num) - k

	for i := range num {
		for top > 0 && stack[top-1] > num[i] && k > 0 {
			top--
			k--
		}
		stack[top] = num[i]
		top++
	}

	i := 0
	for i < digits && stack[i] == '0' {
		i++
	}
	if i == digits {
		return "0"
	}

	return string(stack[i:digits])
}
