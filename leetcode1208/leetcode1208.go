package leetcode1208

func equalSubstring(s string, t string, maxCost int) int {
	i, j := 0, 0
	n := len(s)
	res := 0
	var cost int
	for ; j < n; j++ {
		cost = diff(s[j], t[j])
		maxCost -= cost
		for maxCost < 0 {
			cost = diff(s[i], t[i])
			maxCost += cost
			i++
		}
		if res < j-i+1 {
			res = j - i + 1
		}
	}

	return res
}

func diff(a byte, b byte) int {
	if a > b {
		return int(a - b)
	}
	return int(b - a)
}
