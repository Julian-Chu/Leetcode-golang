package LeetCode_383_RansomNote

func canConstruct(ransomNote string, magazine string) bool {
	if len(ransomNote) > len(magazine) {
		return false
	}
	cnt := make([]int, 26) //only lowercase letter

	for i := 0; i < len(magazine); i++ {
		cnt[magazine[i]-'a']++
	}

	for i := 0; i < len(ransomNote); i++ {
		cnt[ransomNote[i]-'a']--
		if cnt[ransomNote[i]-'a'] < 0 {
			return false
		}
	}

	return true
}

func canConstruct1(ransomNote string, magazine string) bool {
	r := [26]byte{}

	for i := range ransomNote {
		r[ransomNote[i]-'a']++
	}

	for i := range magazine {
		if r[magazine[i]-'a'] != 0 {
			r[magazine[i]-'a']--
		}
	}

	return r == [26]byte{}
}
