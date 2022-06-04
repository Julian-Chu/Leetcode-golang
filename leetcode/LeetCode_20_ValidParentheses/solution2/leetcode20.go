package solution2

func isValid(s string) bool {
	if s == "" {
		return true
	}

	stack := make([]byte, 0, len(s))
	hash := map[byte]byte{'}': '{', ']': '[', ')': '('}

	for i := range s {
		b := s[i]

		if b == '{' || b == '(' || b == '[' {
			stack = append(stack, b)
			continue
		} else if len(stack) > 0 && stack[len(stack)-1] == hash[b] {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}

	return len(stack) == 0
}
