package leetcode150

import "strconv"

func evalRPN(tokens []string) int {
	if len(tokens) == 0 {
		return 0
	}
	numStack := make([]int, 0, len(tokens))

	for _, v := range tokens {
		num, err := strconv.Atoi(v)
		if err == nil {
			numStack = append(numStack, num)
			continue
		}
		v1, v2 := numStack[len(numStack)-2], numStack[len(numStack)-1]
		numStack = numStack[:len(numStack)-2]
		res := 0
		switch v {
		case "+":
			res = v1 + v2
		case "-":
			res = v1 - v2
		case "*":
			res = v1 * v2
		case "/":
			res = v1 / v2
		}
		numStack = append(numStack, res)
	}

	return numStack[0]

}
