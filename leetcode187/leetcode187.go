package leetcode187

func findRepeatedDnaSequences(s string) []string {
	m := make(map[string]int)

	size := 10

	for i := 0; i <= len(s)-size; i++ {
		m[s[i:i+size]]++
	}

	var res []string

	for k, v := range m {
		if v > 1 {
			res = append(res, k)
		}
	}

	return res
}
