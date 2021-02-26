package leetcode227

func calculate(s string) int {
	l := len(s)

	var n1, n2, n3 int
	var opt1, opt2 byte
	opt1 = '+'

	idx := 0

	nextN := func() int {
		n := 0
		for idx < l && s[idx] == ' ' {
			idx++
		}

		for idx < l && '0' <= s[idx] && '9' >= s[idx] {
			n = n*10 + int(s[idx]-'0')
			idx++
		}
		return n
	}

	nextOpt := func() byte {
		opt := byte('+')

		for idx < l && s[idx] == ' ' {
			idx++
		}

		if idx < l {
			opt = s[idx]
			idx++
		}

		return opt
	}

	n2 = nextN()
	for idx < l {
		opt2 = nextOpt()
		n3 = nextN()

		if opt2 == '*' || opt2 == '/' {
			n2 = operate(n2, n3, opt2)
		} else {
			n1 = operate(n1, n2, opt1)
			opt1 = opt2
			n2 = n3
		}
	}

	return operate(n1, n2, opt1)
}

func operate(a, b int, opt byte) int {
	switch opt {
	case '+':
		return a + b
	case '-':
		return a - b
	case '*':
		return a * b
	default:
		return a / b
	}
}
