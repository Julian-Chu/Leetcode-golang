package LeetCode_657_RobotReturnToOrigin

func judgeCircle(moves string) bool {
	cur := [2]int{0, 0}

	m := map[byte][2]int{
		'U': {0, 1},
		'D': {0, -1},
		'L': {-1, 0},
		'R': {1, 0},
	}

	for i := range moves {
		cmd := moves[i]

		offset := m[cmd]
		cur[0] += offset[0]
		cur[1] += offset[1]
	}

	return cur == [2]int{0, 0}
}

func judgeCircle_1(moves string) bool {
	x, y := 0, 0

	for i := range moves {
		switch moves[i] {
		case 'U':
			y++
		case 'D':
			y--
		case 'L':
			x--
		case 'R':
			x++
		}
	}
	return x == 0 && y == 0
}
