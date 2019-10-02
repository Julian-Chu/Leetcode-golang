package leetcode11

func maxArea(height []int) int {
	max := 0
	for i := 0; i < len(height)-1; i++ {
		for j := i + 1; j < len(height); j++ {
			h := height[i]
			if height[i] > height[j] {
				h = height[j]
			}
			v := h * (j - i)
			if v > max {
				max = v
			}
		}
	}
	return max
}
