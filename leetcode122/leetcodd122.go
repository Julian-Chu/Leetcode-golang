package leetcode122

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	sum := 0
	max := 0
	temp := 0

	for i := 1; i < len(prices); i++ {

		if prices[i] < prices[i-1] {
			sum += max
			max = 0
			temp = 0
		} else {
			temp += prices[i] - prices[i-1]
			if max < temp {
				max = temp
			}
		}
		if i == len(prices)-1 {
			sum += temp
		}
	}
	return sum
}
