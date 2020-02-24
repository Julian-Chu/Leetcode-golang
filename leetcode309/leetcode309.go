package leetcode309

func maxProfit(prices []int) int {
	l := len(prices)
	if l == 0 {
		return 0
	}

	buy := make([]int, l+1)
	sell := make([]int, l+1)

	buy[1] = -prices[0]
	for i := 2; i <= l; i++ {
		buy[i] = max(buy[i-1], sell[i-2]-prices[i-1])
		sell[i] = max(sell[i-1], buy[i-1]+prices[i-1])
	}

	return sell[l]

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
