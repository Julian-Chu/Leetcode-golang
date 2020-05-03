package leetcode621

func leastInterval(tasks []byte, n int) int {
	intervals := make([]int, 26)
	taskCount := len(tasks)
	ts := make([]int, 26)

	for _, task := range tasks {
		ts[task-'A']++
	}

	cnt := 0
	for taskCount > 0 {
		maxIdx := -1
		maxCnt := 0
		for i := range ts {
			if ts[i] > maxCnt && intervals[i] <= 0 {
				maxCnt = ts[i]
				maxIdx = i
			}
		}

		for i := range intervals {
			intervals[i]--
		}

		if maxIdx != -1 {
			intervals[maxIdx] = n
			ts[maxIdx]--
			taskCount--
		}
		cnt++
	}

	return cnt

}
