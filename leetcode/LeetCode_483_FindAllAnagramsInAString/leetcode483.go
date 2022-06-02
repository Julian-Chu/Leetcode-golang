package LeetCode_483_FindAllAnagramsInAString

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
	for i := 0; i < len(s); i++ {
		window[s[i]-'a']++
		if i < n-1 {
			continue
		}
		if i >= n {
			window[s[i-n]-'a']--
		}
		if window == target {
			res = append(res, i-n+1)
		}
	}
	return res

}
