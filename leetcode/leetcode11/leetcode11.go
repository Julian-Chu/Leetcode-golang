package leetcode11

func maxArea(height []int) int {
	max := 0
	i, j := 0, len(height)-1

	for i < j {
		h := height[i]
		if height[i] > height[j] {
			h = height[j]
		}

		area := h * (j - i)
		if max < area {
			max = area
		}

		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return max
}
