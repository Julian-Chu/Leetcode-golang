package Leetcode_977_SquaresOfASortedArray

func sortedSquares(nums []int) []int {
	left, right := 0, len(nums)-1
	res := make([]int, len(nums))

	for i := len(nums) - 1; i >= 0; i-- {
		lSquare := nums[left] * nums[left]
		rSquare := nums[right] * nums[right]
		if lSquare > rSquare {
			res[i] = lSquare
			left++
		} else {
			res[i] = rSquare
			right--
		}
	}
	return res
}
