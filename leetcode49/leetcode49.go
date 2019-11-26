package leetcode49

import "sort"

func groupAnagrams(strs []string) [][]string {
	m := make(map[string]int)

	result := make([][]string, 0)

	for _, v := range strs {
		bytes := []byte(v)
		sort.Slice(bytes, func(i, j int) bool {
			if bytes[i] > bytes[j] {
				return true
			} else {
				return false
			}
		})
		sorted := string(bytes)
		if idx, ok := m[sorted]; ok {
			result[idx] = append(result[idx], v)
			continue
		}
		result = append(result, []string{v})
		m[sorted] = len(result) - 1
	}

	return result
}
