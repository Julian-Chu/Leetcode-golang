package LeetCode_1207_UniqueNumberOfOccurrences

func uniqueOccurrences(arr []int) bool {
	count := [2001]int{}

	for _, num := range arr {
		count[num+1000]++
	}

	m := map[int]struct{}{}
	for _, cnt := range count {
		if cnt == 0 {
			continue
		}
		if _, exist := m[cnt]; exist {
			return false
		}
		m[cnt] = struct{}{}
	}

	return true
}
