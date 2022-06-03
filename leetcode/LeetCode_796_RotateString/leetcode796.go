package LeetCode_796_RotateString

func rotateString(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}
	size := len(s)
	for i := 0; i < size; i++ {
		if s[:i] == goal[size-i:] && s[i:] == goal[:size-i] {
			return true
		}
	}
	return false
}
