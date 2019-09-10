package leetcode17

func letterCombinations(digits string) []string {
	res := make([]string, 0)
	for i := 0; i < len(digits); i++ {
		if digits[i] < '2' || digits[i] > '9' {
			continue
		}
		res = appendAlphabetTo(res, digits[i])
		//fmt.Printf("%v\n", res)
		//fmt.Println(len(res))
	}

	return res
}

var pn = map[byte][]string{
	'2': {"a", "b", "c"},
	'3': {"d", "e", "f"},
	'4': {"g", "h", "i"},
	'5': {"j", "k", "l"},
	'6': {"m", "n", "o"},
	'7': {"p", "q", "r", "s"},
	'8': {"t", "u", "v"},
	'9': {"w", "x", "y", "z"},
}

func appendAlphabetTo(input []string, number byte) []string {
	if len(input) == 0 {
		return pn[number]
	}

	res := make([]string, 0)

	for _, i := range input {
		for _, j := range pn[number] {
			res = append(res, i+j)
		}
	}
	return res
}
