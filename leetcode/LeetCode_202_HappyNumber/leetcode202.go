package LeetCode_202_HappyNumber

func isHappy(n int) bool {
	res := make(map[int]bool)

	num := n

	for !res[num] {
		res[num] = true
		next := 0
		for num > 0 {
			root := num % 10
			num /= 10
			next += root * root
		}
		num = next
	}

	return res[1]
}

func isHappy1(n int) bool {
	set := make(map[int]struct{})

	sum := func(k int) int {
		res := 0
		for k > 0 {
			var r int
			r, k = k%10, k/10
			res += r * r
		}

		return res
	}

	for {
		n = sum(n)

		if n == 1 {
			return true
		}

		if _, ok := set[n]; ok {
			return false
		}

		set[n] = struct{}{}
	}
	return false
}
