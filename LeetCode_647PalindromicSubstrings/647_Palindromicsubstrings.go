package LeetCode_647PalindromicSubstrings

func countSubstrings(s string) int {
	cnt := 0
	for i := 0; i < len(s); i++ {
		cnt += calcPalindrome(s, i, i)
		cnt += calcPalindrome(s, i, i+1)
	}
	return cnt

}

func calcPalindrome(s string, left int, right int) int {
	res := 0
	for left >= 0 && right < len(s) && s[left] == s[right] {
		res++
		left--
		right++
	}
	return res
}
