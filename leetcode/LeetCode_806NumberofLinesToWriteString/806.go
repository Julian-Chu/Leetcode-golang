package LeetCode_806NumberofLinesToWriteString

func numberOfLines(widths []int, S string) []int {
	//widthMapper := make(map[rune]int)
	//a := 'a'
	//for i := 0; i < len(widths); i++ {
	//	widthMapper[a+int32(i)] = widths[i]
	//}
	//
	//count := 0
	//sumUnit := 0
	//for i := 0; i < len(S); i++ {
	//	sumUnit += widthMapper[rune(S[i])]
	//	if sumUnit > 100 {
	//		count++
	//		sumUnit = widthMapper[rune(S[i])]
	//	}
	//	fmt.Println(sumUnit)
	//}
	//lastLine := sumUnit
	//if lastLine != 0 {
	//	count++
	//}
	//return []int{count, lastLine}

	res := []int{0, 0}
	if len(S) == 0 {
		return res
	}
	res[0] = 1

	for i := 0; i < len(S); i++ {
		if res[1]+widths[S[i]-'a'] > 100 {
			res[0]++
			res[1] = widths[S[i]-'a']
		} else {
			res[1] += widths[S[i]-'a']
		}
	}
	return res
}
