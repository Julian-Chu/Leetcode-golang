package LeetCode_649_Dota2_Senate

func predictPartyVictory(senate string) string {
	R, D := true, true
	flag := 0
	bs := []byte(senate)
	for R && D {
		R = false
		D = false
		for i := range bs {
			if bs[i] == 'R' {
				if flag < 0 {
					bs[i] = 0
				} else {
					R = true
				}
				flag++
				continue
			}

			if bs[i] == 'D' {
				if flag > 0 {
					bs[i] = 0
				} else {
					D = true
				}
				flag--
				continue
			}
		}
	}
	if R {
		return "Radiant"
	}
	return "Dire"
}
