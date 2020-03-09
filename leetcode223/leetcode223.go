package leetcode223

func computeArea(A int, B int, C int, D int, E int, F int, G int, H int) int {
	type Rec struct {
		top, bottom, left, right int
	}

	getArea := func(r Rec) int {
		area := (r.top - r.bottom) * (r.left - r.right)

		if area < 0 {
			area = -area
		}
		return area
	}
	first := Rec{
		top:    D,
		bottom: B,
		left:   A,
		right:  C,
	}
	areaFirst := getArea(first)
	second := Rec{
		top:    H,
		bottom: F,
		left:   E,
		right:  G,
	}
	areaSecond := getArea(second)
	// check no overlap
	areaOverlap := 1
	if first.top < second.bottom || first.bottom > second.top || first.right < second.left || first.left > second.right {
		areaOverlap = 0
	}

	if areaOverlap != 0 {
		overlap := Rec{
			top:    min(first.top, second.top),
			bottom: max(first.bottom, second.bottom),
			left:   max(first.left, second.left),
			right:  min(first.right, second.right),
		}

		areaOverlap = getArea(overlap)
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
