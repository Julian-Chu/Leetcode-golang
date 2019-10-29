package leetcode121

func maxProfit(prices []int) int {
	n := len(prices)
	max := 0
	for i := range prices {
		for j := i + 1; j < n; j++ {
			profit := prices[j] - prices[i]
			if profit <= 0 {
				continue
			}

			if profit > max {
				max = profit
			}
		}
	}
	return max
}
