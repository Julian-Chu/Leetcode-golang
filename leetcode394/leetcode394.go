package leetcode394

import "strings"

func decodeString(s string) string {
	sb := []byte(s)
	res := []string{}
	brackets := make(map[byte]int)

	leftBracket := byte('[')
	rightBracket := byte(']')
	brackets[leftBracket] = 0
	brackets[rightBracket] = 0
	leftBracketIndx := -1
	rightBracketIndx := -1
	num := 0
	for i := 0; i < len(sb); i++ {
		switch {
		case sb[i] >= '0' && sb[i] <= '9' && brackets[leftBracket] == 0:
			num *= 10
			num += int(sb[i] - '0')
		case sb[i] == leftBracket:
			if brackets[leftBracket] == 0 {
				leftBracketIndx = i
			}
			brackets[leftBracket]++
		case sb[i] == rightBracket:
			brackets[rightBracket]++
			if brackets[rightBracket] != brackets[leftBracket] {
				continue
			}
			rightBracketIndx = i

			var encodedStr = string(sb[leftBracketIndx+1 : rightBracketIndx])
			if brackets[rightBracket] != 1 {
				encodedStr = decodeString(encodedStr)
			}

			for i := 0; i < num; i++ {
				res = append(res, encodedStr)
			}
			num = 0
			brackets[leftBracket] = 0
			brackets[rightBracket] = 0
		default:
			if num != 0 {
				continue
			}
			res = append(res, string(sb[i]))
		}
	}
	return strings.Join(res, "")
}
