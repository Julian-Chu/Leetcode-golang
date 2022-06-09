package LeetCode_17_LetterCombinationsOfAPhoneNumber

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

func letterCombinations_1(digits string) []string {
	if digits == "" {
		return nil
	}

	m := map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}

	var res []string
	var dfs func(digits string, idx int, subset []byte)
	dfs = func(digits string, idx int, subset []byte) {
		if len(subset) == len(digits) {
			res = append(res, string(subset))
			return
		}

		num := digits[idx]
		letters := m[num]
		for i := range letters {
			subset = append(subset, letters[i])
			dfs(digits, idx+1, subset)
			subset = subset[:len(subset)-1]
		}
	}
	dfs(digits, 0, nil)
	return res
}
