package leetcode1023

func camelMatch(queries []string, pattern string) []bool {
	res := make([]bool, len(queries))

	for i, query := range queries {
		res[i] = isMatched(query, pattern)
	}
	return res
}

func isMatched(query string, pattern string) bool {
	i := 0
	n := len(pattern)

	for j := range query {
		if i < n && query[j] == pattern[i] {
			i++
			continue
		}
		if query[j] < 'a' {
			return false
		}
	}

	return i == n
}
