package leetcode150

import "strconv"

func evalRPN(tokens []string) int {
	if len(tokens) == 0 {
		return 0
	}
	numStack := make([]int, 0, len(tokens))

	for _, v := range tokens {
		if num, err := strconv.Atoi(v); err == nil {
			numStack = append(numStack, num)
			continue
		}
		v1, v2 := numStack[len(numStack)-2], numStack[len(numStack)-1]
		numStack = numStack[:len(numStack)-2]
		switch v {
		case "+":
			numStack = append(numStack, v1+v2)
		case "-":
			numStack = append(numStack, v1-v2)
		case "*":
			numStack = append(numStack, v1*v2)
		case "/":
			numStack = append(numStack, v1/v2)
		}
	}

	return numStack[0]

}
