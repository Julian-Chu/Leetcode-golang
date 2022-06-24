package LeetCode_922_SortArrayByParityII

func sortArrayByParityII(nums []int) []int {
	res := make([]int, len(nums))

	oddIdx := 1
	evenIdx := 0
	for _, num := range nums {
		if num&1 == 0 {
			res[evenIdx] = num
			evenIdx += 2
		} else {
			res[oddIdx] = num
			oddIdx += 2
		}
	}
	return res
}

func sortArrayByParityII_1(nums []int) []int {
	oddIdx := 1
	for i := 0; i < len(nums); i += 2 {
		if nums[i]&1 == 1 {
			for nums[oddIdx]&1 != 0 {
				oddIdx += 2
			}
			nums[i], nums[oddIdx] = nums[oddIdx], nums[i]
		}
	}
	return nums
}
