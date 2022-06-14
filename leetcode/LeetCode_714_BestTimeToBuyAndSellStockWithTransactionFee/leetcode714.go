package LeetCode_714_BestTimeToBuyAndSellStockWithTransactionFee

func maxProfit(prices []int, fee int) int {
	n := len(prices)
	res := 0
	minPrice := prices[0]

	for i := 1; i < n; i++ {
		if minPrice > prices[i] {
			minPrice = prices[i]
		}

		if prices[i] >= minPrice && prices[i] <= minPrice+fee {
			continue
		}

		if prices[i] > minPrice+fee {
			res += prices[i] - minPrice - fee
			minPrice = prices[i] - fee // key step
		}
	}
	return res
}
