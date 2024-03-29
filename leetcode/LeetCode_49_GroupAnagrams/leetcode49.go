package LeetCode_49_GroupAnagrams

import "sort"

func groupAnagrams(strs []string) [][]string {
	m := make(map[[26]byte][]string)

	for _, str := range strs {
		arr := [26]byte{}
		for _, ch := range str {
			arr[ch-'a']++
		}

		if _, ok := m[arr]; ok {
			m[arr] = append(m[arr], str)
		} else {
			m[arr] = []string{str}
		}
	}

	res := make([][]string, 0, len(m))
	for _, group := range m {
		res = append(res, group)
	}

	return res
}

// bad performance
func groupAnagrams1(strs []string) [][]string {
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
