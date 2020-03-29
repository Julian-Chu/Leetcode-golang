package leetcode242

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	schars := [26]byte{}
	tchars := [26]byte{}

	for i := range s {
		schars[s[i]-'a']++
	}

	for i := range s {
		tchars[t[i]-'a']++
	}

	for i := range schars {
		if schars[i] != tchars[i] {
			return false
		}
	}

	return true

}
