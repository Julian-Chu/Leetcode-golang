package leetcode309

func maxProfit(prices []int) int {
	l := len(prices)
	if l == 0 {
		return 0
	}
	s0 := 0          //rest after rest, rest after sell
	s1 := -prices[0] // buy after rest, rest after buy
	s2 := -1 << 31   //  sell after rest

	for i := 0; i < len(prices); i++ {
		pre0 := s0
		pre1 := s1
		pre2 := s2

		s0 = max(pre0, pre2)
		s1 = max(pre0-prices[i], pre1)
		s2 = s1 + prices[i]
	}

	return max(s0, s2)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
