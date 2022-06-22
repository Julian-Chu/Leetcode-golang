package LeetCode_941_ValidMountainArray

func validMountainArray(arr []int) bool {
	n := len(arr)
	if n < 3 {
		return false
	}

	l, r := 0, n-1

	for l < n-1 && arr[l] < arr[l+1] {
		l++
	}

	for r > 0 && arr[r-1] > arr[r] {
		r--
	}

	return l == r && l != 0 && r != n-1
}
