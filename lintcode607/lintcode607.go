package lintcode607

var nums = make(map[int]int)

/**
 * @param number: An integer
 * @return: nothing
 */
func add(number int) {
	_, ok := nums[number]
	if !ok {
		nums[number] = 1
		return
	}
	nums[number] += 1
}

/**
 * @param value: An integer
 * @return: Find if there exists any pair of numbers which sum is equal to the value.
 */
func find(value int) bool {
	for num, count := range nums {
		target := value - num
		if target == num {
			return count > 1
		}

		if _, ok := nums[target]; ok {
			return true
		}
	}
	return false
}
