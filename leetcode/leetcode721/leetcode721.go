package leetcode721

import "sort"

func accountsMerge(accounts [][]string) [][]string {
	n := len(accounts)

	parent := make(map[string]string, n)
	owner := make(map[string]string, n)

	for _, account := range accounts {
		for i := 1; i < len(account); i++ {
			parent[account[i]] = account[i]
			owner[account[i]] = account[0]
		}
	}

	for _, account := range accounts {
		r := root(account[1], parent)
		for i := 2; i < len(account); i++ {
			p := root(account[i], parent)
			parent[p] = r
		}
	}
	union := make(map[string][]string, n)
	for email, p := range parent {
		r := root(p, parent)
		union[r] = append(union[r], email)
	}

	res := make([][]string, 0, len(union))

	for p, emails := range union {
		if len(emails) > 1 {
			sort.Strings(emails)
		}
		t := []string{owner[p]}
		t = append(t, emails...)
		res = append(res, t)
	}
	return res
}

func root(e string, parent map[string]string) string {
	if e == parent[e] {
		return e
	}
	return root(parent[e], parent)
}
