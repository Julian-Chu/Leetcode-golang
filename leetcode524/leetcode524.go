package leetcode524

func findLongestWord(s string, d []string) string {
	res := ""
	n := len(s)
	for _, str := range d {
		idx := 0
		i := 0

		size := len(str)
		for idx < size && i < n {
			if str[idx] == s[i] {
				idx++
			}
			i++
		}

		if idx == size {
			switch {
			case len(res) < len(str):
				res = str
			case len(res) == len(str):
				for i := 0; i < len(str); i++ {
					if res[i] < str[i] {
						break
					}

					if res[i] > str[i] {
						res = str
						break
					}
				}
			}

		}
	}

	return res
}
