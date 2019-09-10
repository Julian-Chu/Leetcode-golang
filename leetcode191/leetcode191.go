package leetcode191

func hammingWeight(num uint32) int {
	count := 0
	for num != 0 {
		count++
		num &= num - 1
	}
	return count
}

//func hammingWeight(num uint32) int {
//	if num == 0 {
//		return 0
//	}
//	cnt := 0
//	for i := 0; i < 32; i++ {
//		if num&1 == 1 {
//			cnt++
//		}
//		num = num >> 1
//	}
//	return cnt
//}
