package leetcode415

func addStrings(num1 string, num2 string) string {
	if num1 == "0" && num2 == "0" {
		return "0"
	}
	l1 := len(num1)
	l2 := len(num2)

	b1 := []byte(num1)
	b2 := []byte(num2)

	maxLen := l1
	if l2 > l1 {
		maxLen = l2
	}

	res := make([]byte, maxLen+1)
	carry := byte(0)
	for i, j, k := l1-1, l2-1, maxLen; k >= 0; {
		var a byte
		var b byte
		if i >= 0 {
			a = b1[i] - '0'
			i--
		}

		if j >= 0 {
			b = b2[j] - '0'
			j--
		}
		sum := a + b + carry
		if sum > 9 {
			res[k] = sum%10 + '0'
			carry = 1
		} else {
			res[k] = sum + '0'
			carry = 0
		}
		k--
	}

	if res[0] == '0' {
		res = res[1:]
	}

	return string(res)
}
