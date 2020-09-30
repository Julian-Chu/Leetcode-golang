package leetcode721

import "sort"

func accountsMerge_1(accounts [][]string) [][]string {
	emails := make(map[string]int)
	accts := make([]int, len(accounts))

	for i := 1; i < len(accounts); i++ {
		accts[i] = i
	}
	for acctNum, lst := range accounts {
		for _, email := range lst[1:] {
			if v, ok := emails[email]; ok {
				parent := getParent(acctNum, accts)
				accts[parent] = getParent(v, accts)
			} else {
				emails[email] = getParent(acctNum, accts)
			}
		}
	}
	mp := make(map[int][]string)
	for i, v := range accts {
		if i == getParent(v, accts) {
			mp[i] = []string{accounts[i][0]}
		}
	}

	for email, acct := range emails {
		parent := getParent(acct, accts)
		mp[parent] = append(mp[parent], email)
	}
	res := [][]string{}
	for _, v := range mp {
		sort.Strings(v)
		res = append(res, v)
	}
	return res
}

func getParent(num int, accts []int) int {
	if accts[num] != num {
		accts[num] = getParent(accts[num], accts)
	}
	return accts[num]
}
