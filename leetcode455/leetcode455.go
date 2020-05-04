package leetcode455

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
