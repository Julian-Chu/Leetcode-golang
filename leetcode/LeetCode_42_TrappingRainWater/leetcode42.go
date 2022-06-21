package LeetCode_42_TrappingRainWater

func trap_twoPointers(height []int) int {
	sum := 0
	n := len(height)
	for i := 0; i < n; i++ {
		if i == 0 || i == n-1 {
			continue
		}

		rHeight := height[i]
		lHeight := height[i]

		for r := i + 1; r < n; r++ {
			if rHeight < height[r] {
				rHeight = height[r]
			}
		}

		for l := i - 1; l >= 0; l-- {
			if lHeight < height[l] {
				lHeight = height[l]
			}
		}
		h := min(lHeight, rHeight) - height[i]
		if h > 0 {
			sum += h
		}
	}
	return sum
}

func trap_dp(height []int) int {
	n := len(height)
	if n <= 2 {
		return 0
	}

	maxLeft := make([]int, n)
	maxRight := make([]int, n)

	maxLeft[0] = height[0]
	for i := 1; i < n; i++ {
		maxLeft[i] = max(height[i], maxLeft[i-1])
	}

	maxRight[n-1] = height[n-1]
	for i := n - 2; i >= 0; i-- {
		maxRight[i] = max(height[i], maxRight[i+1])
	}
	sum := 0

	for i := 0; i < n; i++ {
		water := min(maxLeft[i], maxRight[i]) - height[i]
		if water > 0 {
			sum += water
		}
	}
	return sum
}

func trap_monotonicStack(height []int) int {
	n := len(height)
	stack := make([]int, 0, n)
	stack = append(stack, 0)

	sum := 0

	for i := 1; i < n; i++ {
		for len(stack) > 0 && height[i] > height[stack[len(stack)-1]] {
			mid := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) > 0 {
				h := min(height[stack[len(stack)-1]], height[i]) - height[mid]
				w := i - stack[len(stack)-1] - 1
				sum += h * w
			}
		}
		stack = append(stack, i)
	}
	return sum
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
