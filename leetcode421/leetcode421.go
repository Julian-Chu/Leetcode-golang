package leetcode421

func findMaximumXOR(nums []int) int {
	var mask, max int

	for i := 31; i >= 0; i-- {
		mask |= 1 << uint32(i)

		nMap := make(map[int]bool)
		for _, num := range nums {
			nMap[mask&num] = true
		}
		tmp := max | 1<<uint32(i)
		for key := range nMap {
			if nMap[tmp^key] {
				max = tmp
				break
			}
		}
	}
	return max
}
