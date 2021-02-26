package leetcode75

import "fmt"

func sortColors(nums []int) []int {
	quicksort(nums)

	fmt.Println(nums)
	return nums
}

func quicksort(nums []int) {
	if len(nums) == 0 || len(nums) == 1 {
		return
	}
	mid := nums[0]

	l, r := 0, len(nums)-1
	for l < r {
		for nums[r] >= mid && l < r {
			r--
		}
		for nums[l] < mid && l < r {
			l++
		}
		nums[l], nums[r] = nums[r], nums[l]
	}
	quicksort(nums[:l+1])
	quicksort(nums[l+1:])

}
