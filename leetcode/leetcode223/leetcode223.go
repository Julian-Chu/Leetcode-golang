package leetcode223

func computeArea(A int, B int, C int, D int, E int, F int, G int, H int) int {

	getArea := func(A, B, C, D int) int {
		area := (A - C) * (B - D)

		if area < 0 {
			area = -area
		}
		return area
	}
	areaFirst := getArea(A, B, C, D)
	areaSecond := getArea(E, F, G, H)
	// check no overlap
	areaOverlap := 1
	if A > G || C < E || D < F || B > H {
		areaOverlap = 0
	}

	if areaOverlap != 0 {
		areaOverlap = getArea(max(A, E), max(B, F), min(C, G), min(D, H))
	}

	return areaFirst - areaOverlap + areaSecond
}

func min(a, b int) int {
	if a > b {
		return b
	}

	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
