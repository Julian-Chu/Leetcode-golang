package LeetCode_1047_RemoveAllAdjacentDuplicatesInString

func removeDuplicates(s string) string {
	stack := make([]byte, 0, len(s))

	for i := range s {
		b := s[i]

		if len(stack) > 0 && b == stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
			continue
		}

		stack = append(stack, b)
	}

	return string(stack)
}

func removeDuplicates_TwoPointers(s string) string {
	slow := 0
	bs := []byte(s)

	for i := range bs {
		bs[slow] = bs[i]

		if slow > 0 && bs[slow] == bs[slow-1] {
			slow--
		} else {
			slow++
		}
	}
	return string(bs[:slow])
}
