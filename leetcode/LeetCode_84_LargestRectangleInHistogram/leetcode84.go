package LeetCode_84_LargestRectangleInHistogram

func largestRectangleArea(heights []int) int {
	res := 0

	n := len(heights)
	for i := 0; i < n; i++ {
		h := heights[i]
		for j := i; j < n; j++ {
			if heights[j] < h {
				h = heights[j]
			}
			w := j - i + 1
			area := h * w
			if area > res {
				res = area
			}
		}
	}
	return res
}

func largestRectangleArea_dp(heights []int) int {
	n := len(heights)
	minLeftIdx := make([]int, n)
	minRightIdx := make([]int, n)

	minLeftIdx[0] = -1

	for i := 0; i < n; i++ {
		prev := i - 1
		for prev >= 0 && heights[prev] >= heights[i] {
			prev = minLeftIdx[prev]
		}
		minLeftIdx[i] = prev
	}

	minRightIdx[n-1] = n
	for j := n - 2; j >= 0; j-- {
		next := j + 1
		for next < n && heights[next] >= heights[j] {
			next = minRightIdx[next]
		}
		minRightIdx[j] = next
	}

	res := 0

	for i := 0; i < n; i++ {
		sum := heights[i] * (minRightIdx[i] - minLeftIdx[i] - 1)
		res = max(sum, res)
	}
	return res
}

func largestRectangleArea_montonicstack(heights []int) int {
	heights = append(heights, 0)
	heights = append([]int{0}, heights...)
	n := len(heights)
	stack := make([]int, 1, n)
	stack[0] = 0
	res := 0

	for i := 1; i < n; i++ {
		top := stack[len(stack)-1]
		if heights[i] > heights[top] {
			stack = append(stack, i)
		} else if heights[i] == heights[top] {
			stack = stack[:len(stack)-1]
			stack = append(stack, i)
		} else {
			for heights[i] < heights[top] {
				h := heights[top]
				stack = stack[:len(stack)-1]
				left := stack[len(stack)-1]
				right := i
				w := right - left - 1
				if res < h*w {
					res = h * w
				}
				top = stack[len(stack)-1]
			}
			stack = append(stack, i)
		}
	}
	return res
}
func largestRectangleArea_dp2(heights []int) int {
	heights = append(heights, 0)
	heights = append([]int{0}, heights...)
	n := len(heights)
	stack := []int{0}
	res := 0

	for i := 1; i < n; i++ {
		top := stack[len(stack)-1]
		for heights[i] < heights[top] {
			h := heights[top]
			stack = stack[:len(stack)-1]
			left := stack[len(stack)-1]
			right := i
			w := right - left - 1
			if res < h*w {
				res = h * w
			}
			top = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}
	return res
}
func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
