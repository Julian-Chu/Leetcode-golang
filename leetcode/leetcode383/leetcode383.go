package leetcode383

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
