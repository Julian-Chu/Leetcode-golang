package LeetCode_27_RemoveElement

func removeElement(nums []int, val int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		for left <= right && nums[right] == val {
			right--
		}

		for left <= right && nums[left] != val {
			left++
		}

		if left <= right {
			nums[left], nums[right] = nums[right], nums[left]
			left++
			right--
		}
	}

	return left
}

//func removeElement(nums []int, val int) int {
//	i, j := 0, len(nums)-1
//
//	for {
//		for i < len(nums) && nums[i] != val {
//			i++
//		}
//
//		for j >= 0 && nums[j] == val {
//			j--
//		}
//		if j < i {
//			break
//		}
//		fmt.Println(i, j)
//		nums[i], nums[j] = nums[j], nums[i]
//	}
//
//	return i
//}
