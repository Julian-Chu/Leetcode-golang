package leetcode447

func numberOfBoomerangs(points [][]int) int {
	res := 0
	n := len(points)
	if n < 3 {
		return 0
	}

	calcDistSquare := func(p1, p2 []int) int {
		x := p1[0] - p2[0]
		y := p1[1] - p2[1]

		return x*x + y*y
	}

	for i := 0; i < n; i++ {
		distMap := make(map[int]int, n)
		for j := 0; j < n; j++ {
			if j == i {
				continue
			}
			dist := calcDistSquare(points[i], points[j])
			distMap[dist]++
		}

		for _, cnt := range distMap {
			res += cnt * (cnt - 1)
		}
	}

	return res
}
