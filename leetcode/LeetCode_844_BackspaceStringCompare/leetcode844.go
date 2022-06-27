package LeetCode_844_BackspaceStringCompare

func backspaceCompare(s string, t string) bool {
	return helper(s) == helper(t)
}

func helper(str string) string {
	bs := make([]byte, 0)
	backspaces := 0
	for i := len(str) - 1; i >= 0; {
		if str[i] == '#' {
			backspaces++
			i--
			continue
		}

		if backspaces > 0 {
			i--
			backspaces--
			continue
		}

		bs = append(bs, str[i])
		i--
	}
	return string(bs)
}
