package leetcode832

func flipAndInvertImage(A [][]int) [][]int {
	for _, arr := range A {
		for i, j := 0, len(A)-1; i <= j; i, j = i+1, j-1 {
			if i != j {
				arr[i], arr[j] = arr[j]^1, arr[i]^1
			} else {
				arr[i] = arr[i] ^ 1
			}
		}
		//for i := 0; i < len(arr); i++ {
		//	arr[i] = arr[i] ^ 1
		//}
	}
	return A
}
