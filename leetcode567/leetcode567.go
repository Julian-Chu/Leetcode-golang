package leetcode567

func checkInclusion(s1 string, s2 string) bool {
	arr1 := [26]int{}
	n := len(s1)
	for _, ch := range s1 {
		arr1[ch-'a']++
	}

	arr2 := [26]int{}

	for i := 0; i < len(s2)-n+1; i++ {
		if i == 0 {
			for _, ch := range s2[:n] {
				arr2[ch-'a']++
			}
		} else {
			arr2[s2[i-1]-'a']--
			arr2[s2[i+n-1]-'a']++
		}

		if arr1 == arr2 {
			return true
		}
	}

	return false

}
