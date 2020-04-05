package leetcode438

func findAnagrams(s string, p string) []int {
	n := len(p)
	if len(s) == 0 || len(s) < n {
		return []int{}
	}

	target := [26]int{}

	for i := range p {
		target[p[i]-'a']++
	}

	window := [26]int{}
	res := make([]int, 0)
	for i := 0; i <= len(s)-n; i++ {
		if i == 0 {
			subStr := s[i : i+n]
			for j := range subStr {
				window[subStr[j]-'a']++
			}
		} else {
			window[s[i-1]-'a']--
			window[s[i+n-1]-'a']++
		}
		if target == window {
			res = append(res, i)
		}
	}
	return res

}
