package leetcode43

func multiply(num1 string, num2 string) string {

	if num1[0] == '0' || num2[0] == '0' {
		return "0"
	}
	temp := make([]int, len(num1)+len(num2))

	for i := 0; i < len(num1); i++ {
		for j := 0; j < len(num2); j++ {
			temp[i+j+1] += int((num1[i] - '0') * (num2[j] - '0'))
		}
	}

	for i := len(temp) - 1; i >= 0; i-- {
		if temp[i] >= 10 {
			temp[i-1] += temp[i] / 10
			temp[i] %= 10
		}
	}
	if temp[0] == 0 {
		temp = temp[1:]
	}
	res := make([]byte, len(temp))
	for i := range temp {
		res[i] = byte(temp[i]) + '0'
	}

	return string(res)
}
