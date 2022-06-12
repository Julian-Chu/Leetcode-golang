package LeetCode_455_AssignCookies

func findContentChildren(g []int, s []int) int {
	qsort(g)
	qsort(s)
	j := 0
	cnt := 0
	for i := 0; i < len(s) && j < len(g); i++ {
		if s[i] >= g[j] {
			j++
			cnt++
		}
	}
	return cnt
}

func qsort(data []int) {
	if len(data) <= 1 {
		return
	}

	head := 0
	tail := len(data) - 1
	mid := data[0]

	for i := 1; i <= tail; {
		if data[i] > mid {
			data[i], data[tail] = data[tail], data[i]
			tail--
		} else {
			data[i], data[head] = data[head], data[i]
			head++
			i++
		}
	}
	qsort(data[:head])
	qsort(data[head+1:])
}

func findContentChildren_1(g []int, s []int) int {
	qsort(g)
	qsort(s)

	res := 0

	p, q := 0, 0

	for p < len(g) && q < len(s) {
		if g[p] <= s[q] {
			res++
			p++
			q++
			continue
		}

		q++
	}

	return res
}

func qsort_1(s []int) {
	if len(s) == 0 {
		return
	}
	l, r := 0, len(s)-1

	mid := s[(l + (r-l)>>1)]

	for l <= r {
		for l <= r && s[l] < mid {
			l++
		}
		for l <= r && s[r] > mid {
			r--
		}

		if l <= r {
			s[l], s[r] = s[r], s[l]
			l++
			r--
		}
	}

	qsort(s[0 : r+1])
	qsort(s[l:len(s)])
}
