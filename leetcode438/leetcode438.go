package leetcode438

func findAnagrams(s string, p string) []int {
	if len(s) == 0 {
		return []int{}
	}

	res := make([]int, 0)
	m := make(map[byte]int)

	n := len(p)
	var xorRes byte = 0
	sumRes := 0
	for i, _ := range p {
		xorRes ^= p[i]
		sumRes += int(p[i])
		m[p[i]]++
	}

	xor := uint8(0)
	sum := 0
	for i := 0; i <= len(s)-n; i++ {
		if i == 0 {
			for j := 0; j < n; j++ {
				xor ^= s[j]
				sum += int(s[j])
			}
		} else {
			xor ^= s[i-1]
			xor ^= s[i+n-1]
			sum -= int(s[i-1])
			sum += int(s[i+n-1])
		}

		if xor == xorRes && sum == sumRes {
			msub := make(map[byte]int)
			subStr := s[i : i+n]
			for j := range subStr {
				msub[subStr[j]]++
			}

			flag := true
			for key, v := range msub {
				cnt, ok := m[key]
				if !ok || (ok && cnt != v) {
					flag = false
					break
				}
			}
			if flag {
				res = append(res, i)
			}
		}

	}

	return res
}
