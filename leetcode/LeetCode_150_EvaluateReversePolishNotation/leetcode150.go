package LeetCode_150_EvaluateReversePolishNotation

import "strconv"

func evalRPN(tokens []string) int {
	stack := make([]int, 0, len(tokens))
	for _, str := range tokens {
		num, err := strconv.Atoi(str)
		if err != nil {
			r := stack[len(stack)-1]
			l := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			res := 0
			switch str {
			case "+":
				res = l + r
			case "-":
				res = l - r
			case "/":
				res = l / r
			case "*":
				res = l * r
			}

			stack = append(stack, res)
		} else {
			stack = append(stack, num)
		}
	}

	return stack[0]
}
