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

func helper_1(s string) string {
	bs := []byte{}
	for i := 0; i < len(s); i++ {
		if s[i] == '#' {
			if len(bs) > 0 {
				bs = bs[:len(bs)-1]
			}
		} else {
			bs = append(bs, s[i])
		}
	}

	return string(bs)
}

func backspaceCompare_2(s string, t string) bool {
	i, j := len(s)-1, len(t)-1

	// #abc, abc
	for i >= 0 || j >= 0 {
		skip := 0
		for i >= 0 {
			if s[i] == '#' {
				skip++
				i--
			} else if skip > 0 {
				skip--
				i--
			} else {
				break
			}
		}

		skip = 0
		for j >= 0 {
			if t[j] == '#' {
				skip++
				j--
			} else if skip > 0 {
				skip--
				j--
			} else {
				break
			}
		}

		// important to check i and j >= 0
		if i >= 0 && j >= 0 && s[i] != t[j] {
			return false
		}
		// some ch remains
		if (i >= 0 && j < 0) || (i < 0 && j >= 0) {
			return false
		}
		i--
		j--
	}
	return true
}
