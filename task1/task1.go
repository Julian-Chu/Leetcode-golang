package task1

func Solution(A []int) bool {
	// write your code in Go 1.4
	n := len(A)
	if n < 5 || n > 100000 {
		return false
	}
	sumOfA := 0

	for i := range A {
		sumOfA += A[i]
	}

	tWorker1 := 0

	for i := 0; i < n-4; i++ {
		tWorker1 += A[i]
		drop1 := i + 1
		tWorker2 := 0
		for j := drop1 + 1; j < n-2; j++ {
			tWorker2 += A[j]
			drop2 := j + 1

			if tWorker1 == tWorker2 {
				tWorker3 := sumOfA - tWorker1 - tWorker2 - A[drop1] - A[drop2]

				if tWorker3 == tWorker1 {
					return true
				}
			}

			if tWorker1 < tWorker2 {
				continue
			}

		}
	}
	return false
}
