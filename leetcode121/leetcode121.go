package leetcode121

func maxProfit(prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}
	max := 0
	min := prices[0]
	for i := range prices {
		if max < prices[i]-min {
			max = prices[i] - min
		}
		if min > prices[i] {
			min = prices[i]
		}
	}
	return max
}
