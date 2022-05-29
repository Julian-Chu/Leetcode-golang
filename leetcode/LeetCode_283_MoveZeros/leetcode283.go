package LeetCode_283_MoveZeros

func moveZeroes(nums []int) {
	replace := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[replace] = nums[i]
			replace++
		}
	}

	for i := replace; i < len(nums); i++ {
		nums[i] = 0
	}
}

//func moveZeroes(nums []int)  {
//	if len(nums)<=1 {
//		return
//	}
//	fast, slow:= 0, 0
//
//	for fast <len(nums){
//		if nums[fast] != 0{
//			nums[slow] = nums[fast]
//			slow++
//		}
//		fast++
//	}
//
//	for slow<len(nums){
//		nums[slow] = 0
//		slow++
//	}
//}
