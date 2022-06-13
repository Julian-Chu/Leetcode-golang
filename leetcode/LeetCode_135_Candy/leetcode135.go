package LeetCode_135_Candy

func candy(ratings []int) int {
	n := len(ratings)
	candies := make([]int, n)

	for i := 0; i < n; i++ {
		if i == 0 {
			candies[0] = 1
		} else {
			if ratings[i-1] >= ratings[i] {
				candies[i] = 1
				if candies[i-1] > 1 {
					continue
				}

				j := i - 1
				for j >= 0 && ratings[j] > ratings[j+1] && candies[j] == candies[j+1] {
					candies[j] += 1
					j--
				}
				continue
			}
			candies[i] = candies[i-1] + 1
		}
	}

	sum := 0
	for _, v := range candies {
		sum += v
	}
	return sum
}

func candy_1(ratings []int) int {
	n := len(ratings)
	candies := make([]int, n)
	for i := range candies {
		candies[i] = 1
	}

	for i := 1; i < n; i++ {
		if ratings[i] > ratings[i-1] {
			candies[i] = candies[i-1] + 1
		}
	}

	for i := n - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			candies[i] = max(candies[i], candies[i+1]+1)
		}
	}
	sum := 0
	for _, v := range candies {
		sum += v
	}
	return sum
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
