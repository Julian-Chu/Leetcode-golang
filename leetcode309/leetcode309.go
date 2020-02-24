package leetcode309

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	s0 := 0          // sell->rest or rest->rest
	s1 := -prices[0] // rest->buy or buy->rest
	s2 := -1 << 31   // rest->sell

	for i := 1; i < len(prices); i++ {
		pre0 := s0
		pre1 := s1
		pre2 := s2
		s0 = max(pre0, pre2)
		s1 = max(pre0-prices[i], pre1)
		s2 = pre1 + prices[i]
	}
	return max(s0, s2)
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
