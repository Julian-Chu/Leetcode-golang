package LeetCode_242_ValidAnagram

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	schars := [26]int{}

	for i := range s {
		schars[s[i]-'a']++
		schars[t[i]-'a']--
	}

	return schars == [26]int{}
}

func isAnagram1(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sArr := [128]int{}
	tArr := [128]int{}

	for i := range s {
		sArr[s[i]]++
		tArr[t[i]]++
	}

	return sArr == tArr

}
