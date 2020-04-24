package leetcode406

import "sort"

func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		if people[i][1] < people[j][1] {
			return true
		}

		if people[i][1] == people[j][1] && people[i][0] < people[j][0] {
			return true
		}

		return false
	})

	temp := make([][]int, 0, len(people))
	res := make([][]int, 0, len(people))
	for len(people) > 0 || len(temp) > 0 {
		if len(res) == 0 {
			res = append(res, people[0])
			people = people[1:]
			continue
		}
		var p []int
		if len(people) > 0 {
			p = people[0]
			people = people[1:]
		} else {
			p = temp[0]
			temp = temp[1:]
		}
		v := p[0]
		k := p[1]

		i := 0

		for ; i < len(res); i++ {
			if k == 0 {
				if res[i][0] < v {
					continue
				}

				if res[i][0] >= v {
					break
				}
			}
			if res[i][0] >= v {
				k--
				continue
			}
		}

		if k > 0 {
			temp = append(temp, p)
		} else {
			res = append(res, p)
			copy(res[i+1:], res[i:])
			res[i] = p
		}
	}

	return res
}
