package leetcode990

func equationsPossible(equations []string) bool {
	roots := [26]uint8{}

	for i := range roots {
		roots[i] = uint8(i)
	}

	for _, eq := range equations {
		if eq[1] == '=' {
			v1 := eq[0] - 'a'
			v2 := eq[3] - 'a'
			p1 := getRoot(v1, roots)
			p2 := getRoot(v2, roots)
			roots[p1] = p2
		}
	}

	for _, eq := range equations {
		if eq[1] == '!' {
			v1 := eq[0] - 'a'
			v2 := eq[3] - 'a'
			p1 := getRoot(v1, roots)
			p2 := getRoot(v2, roots)
			if p1 == p2 {
				return false
			}
		}
	}
	return true
}

func getRoot(v uint8, roots [26]uint8) uint8 {
	if v == roots[v] {
		return v
	}
	return getRoot(roots[v], roots)
}
