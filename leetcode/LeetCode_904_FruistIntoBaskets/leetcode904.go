package LeetCode_904_FruistIntoBaskets

func totalFruit(fruits []int) int {
	if len(fruits) <= 1 {
		return len(fruits)
	}
	lastFruit, lastFruitCount := -1, 0
	firstFruit := 0
	count := 0
	res := 0
	for _, t := range fruits {
		count++
		if t != firstFruit && t != lastFruit {
			count = lastFruitCount + 1
		}

		if res < count {
			res = count
		}

		lastFruitCount++
		if t != lastFruit {
			firstFruit, lastFruit = lastFruit, t
			lastFruitCount = 1
		}
	}
	return res
}

//func totalFruit(fruits []int) int {
//	if len(fruits) <= 1 {
//		return len(fruits)
//	}
//	slow, fast := 0, 0
//	sum, res := 0, 0
//	set := make(map[int]int)
//
//	for slow <= fast && fast < len(fruits) {
//		fruit := fruits[fast]
//		_, exist := set[fruit]
//
//		if exist || len(set) < 2 {
//			if exist {
//				set[fruit]++
//			} else {
//				set[fruit] = 1
//			}
//			sum++
//			if res < sum {
//				res = sum
//			}
//			fast++
//			continue
//		}
//
//		set[fruits[slow]]--
//		if set[fruits[slow]] == 0{
//			delete(set, fruits[slow])
//		}
//
//		sum--
//		slow++
//	}
//	return res
//}
