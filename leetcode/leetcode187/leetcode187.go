package leetcode187

func findRepeatedDnaSequences(s string) []string {
	m := make(map[string]int)

	size := 10
	var res []string
	for i := 0; i <= len(s)-size; i++ {
		str := s[i : i+size]
		if m[str] == 1 {
			res = append(res, str)
		}
		m[str]++
	}

	return res
}
