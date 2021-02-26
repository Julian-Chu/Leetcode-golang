package leetcode721

import "sort"

func accountsMerge_1(accounts [][]string) [][]string {
	n := len(accounts)
	parent := make([]int, n)
	emails := make(map[string]int)

	for i := range accounts {
		parent[i] = i
	}

	for acctIdx, account := range accounts {
		for i := 1; i < len(account); i++ {
			if v, ok := emails[account[i]]; ok {
				p := getParent(acctIdx, parent)
				parent[p] = getParent(v, parent)
			} else {
				emails[account[i]] = getParent(acctIdx, parent)
			}
		}
	}
	mp := make(map[int][]string)

	for email, i := range emails {
		p := getParent(i, parent)
		mp[p] = append(mp[p], email)
	}

	res := make([][]string, 0, len(mp))
	for idx, emails := range mp {

		t := make([]string, len(emails)+1)
		t[0] = accounts[idx][0]
		copy(t[1:], emails)
		sort.Strings(t)
		res = append(res, t)
	}

	return res
}

func getParent(idx int, parent []int) int {
	if idx == parent[idx] {
		return idx
	}

	return getParent(parent[idx], parent)
}
