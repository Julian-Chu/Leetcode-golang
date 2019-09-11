package leetcode22

func generateParenthesis(n int) []string {
	res := make([]string, 0)

	cnt := n
	for cnt > 0 {
		if len(res) == 0 {
			res = append(res, "()")
			cnt--
			continue
		}

		temp := make([]string, 0)
		for k, v := range res {
			temp = append(temp, "("+v+")")
			temp = append(temp, v+"()")
			if k != len(res)-1 {
				temp = append(temp, "()"+v)
			}
		}
		cnt--
		res = temp
	}
	//sort.Slice(res, func(i, j int) bool {
	//	for k := 0; k < len(res[i]); k++ {
	//		if res[i][k] < res[j][k] {
	//			return false
	//		}
	//	}
	//	return true
	//})
	return res

}
