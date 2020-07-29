package leetcode477

func totalHammingDistance(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	xor := make([]int, 1, len(nums)-1)

	xor[0] = nums[0]
	sum := 0
	for i := 1; i < len(nums); i++ {
		for j := range xor {
			sum += HammingDistance(xor[j] ^ nums[i])
		}
		xor = append(xor, nums[i])
	}
	return sum
}

func HammingDistance(num int) int {
	cnt := 0
	for num > 0 {
		cnt += num & 1
		num >>= 1
	}
	return cnt
}
