package leetcode242

func isAnagram_1(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	schars := [26]byte{}
	tchars := [26]byte{}

	for i := range s {
		schars[s[i]-'a']++
	}

	for i := range t {
		tchars[t[i]-'a']++
	}

	return schars == tchars
}
