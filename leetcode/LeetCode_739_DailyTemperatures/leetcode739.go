package LeetCode_739_DailyTemperatures

func dailyTemperatures(temperatures []int) []int {
	stack := make([]int, 0, len(temperatures))
	res := make([]int, len(temperatures))

	stack = append(stack, 0)
	for i := 1; i < len(temperatures); i++ {
		top := stack[len(stack)-1]
		if temperatures[i] < temperatures[top] {
			stack = append(stack, i)
		} else if temperatures[i] == temperatures[top] {
			stack = append(stack, i)
		} else {
			for len(stack) != 0 && temperatures[i] > temperatures[top] {
				res[top] = i - top
				stack = stack[:len(stack)-1]
				if len(stack) != 0 {
					top = stack[len(stack)-1]
				}
			}

			stack = append(stack, i)
		}
	}
	return res
}

func dailyTemperatures_1(temperatures []int) []int {
	stack := make([]int, 0)
	res := make([]int, len(temperatures))

	for i, v := range temperatures {
		for len(stack) != 0 && v > temperatures[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res[top] = i - top
		}
		stack = append(stack, i)
	}
	return res
}

func dailyTemperatures_brute(temperatures []int) []int {
	res := make([]int, len(temperatures))

	n := len(temperatures)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if temperatures[j] > temperatures[i] {
				res[i] = j - i
				break
			}
		}
	}

	return res
}
