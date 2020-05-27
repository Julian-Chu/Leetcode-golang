package leetcode424

func characterReplacement(s string, k int) int {
	var max, l int //max : max frequency of uppercase English letter
	count := [128]int{}

	for r := 0; r < len(s); r++ {
		count[s[r]]++
		max = Max(max, count[s[r]])

		if r-l+1-max > k {
			count[s[l]]--
			l++
		}
	}

	return len(s) - l
}

func Max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
