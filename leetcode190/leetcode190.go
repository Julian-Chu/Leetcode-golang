package leetcode190

func reverseBits(num uint32) uint32 {
	if num == 0 {
		return 0
	}

	var res uint32 = 0
	for i := 0; i < 32; i++ {
		res = res << 1
		if num&1 == 1 {
			res++
		}
		num = num >> 1
	}
	return res
}
