package leetcode260

func singleNumber(nums []int) []int {

	xor := 0
	for _, num := range nums {
		xor ^= num
	}
	diff := xor & -xor

	res := make([]int, 2)

	for _, num := range nums {
		if num&diff != 0 {
			res[0] ^= num
		} else {
			res[1] ^= num
		}
	}
	return res
}
