package leetcode299

import "fmt"

func getHint(secret string, guess string) string {
	bs1 := []byte(secret)
	bs2 := []byte(guess)
	m := make(map[byte]int)
	cntA := 0
	for i := range bs1 {
		m[bs1[i]]++
	}

	for i, v := range bs2 {
		cnt, ok := m[v]
		if !ok {
			continue
		}

		if bs1[i] == v {
			cntA++
		}

		if cnt > 0 {
			m[v]--
		}
	}

	cntB := len(bs1) - cntA

	for _, v := range m {
		cntB -= v
	}

	return fmt.Sprintf("%vA%vB", cntA, cntB)

}
