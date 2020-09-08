package leetcode659

func isPossible(nums []int) bool {
	size := len(nums)
	count := make(map[int]int, size)
	next := make(map[int]int, size)

	for i := range nums {
		count[nums[i]]++
	}

	for _, n := range nums {
		if count[n] == 0 {
			continue
		}

		count[n]--

		if next[n] > 0 {
			next[n]--
			next[n+1]++
		} else if count[n+1] > 0 && count[n+2] > 0 {
			count[n+1]--
			count[n+2]--
			next[n+3]++
		} else {
			return false
		}
	}

	return true
}

func isPossible2(nums []int) bool {
	c1, c2, c3 := 1, 0, 0
	p1, p2, p3 := 0, 0, 0
	size := len(nums)

	for i := 1; i < size; i++ {
		if nums[i] == nums[i-1] {
			switch {
			case p1 > 0:
				p1--
				c2++
			case p2 > 0:
				p2--
				c3++
			case p3 > 0:
				p3--
				c3++
			default:
				c1++
			}
		} else if nums[i] == nums[i-1]+1 {
			// subsequences remain
			if p1 > 0 || p2 > 0 {
				return false
			}
			p1, p2, p3 = c1, c2, c3
			c1, c2, c3 = 0, 0, 0
			switch {
			case p1 > 0:
				p1--
				c2++
			case p2 > 0:
				p2--
				c3++
			case p3 > 0:
				p3--
				c3++
			default:
				c1++
			}
		} else {
			if c1 > 0 || c2 > 0 {
				return false
			}
			p1, p2, p3 = 0, 0, 0
			c1, c2, c3 = 1, 0, 0
		}
	}
	return c1 == 0 && c2 == 0
}
