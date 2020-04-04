package leetcode389

func findTheDifference(s string, t string) byte {
	res := int32(0)
	for _, v := range s {
		res ^= v
	}

	for _, v := range t {
		res ^= v // xor associativity
	}

	return byte(res)
}
