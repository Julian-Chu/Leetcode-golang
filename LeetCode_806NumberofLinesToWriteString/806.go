package LeetCode_806NumberofLinesToWriteString

import "fmt"

func numberOfLines(widths []int, S string) []int {
	widthMapper := make(map[rune]int)
	a := 'a'
	for i := 0; i < len(widths); i++ {
		widthMapper[a+int32(i)] = widths[i]
	}

	count := 0
	sumUnit := 0
	for i := 0; i < len(S); i++ {
		sumUnit += widthMapper[rune(S[i])]
		if sumUnit > 100 {
			count++
			sumUnit = widthMapper[rune(S[i])]
		}
		fmt.Println(sumUnit)
	}
	lastLine := sumUnit
	if lastLine != 0 {
		count++
	}
	//for key, val := range widthMapper {
	//	fmt.Println(string(key), " ", val)
	//}
	return []int{count, lastLine}
}
