package leetcode648

import (
	"sort"
	"strings"
)

func replaceWords_2(dictionary []string, sentence string) string {
	sort.Strings(dictionary)
	ss := strings.Split(sentence, " ")
	res := make([]string, len(ss))
	for i := range ss {
		for _, root := range dictionary {
			if strings.HasPrefix(ss[i], root) {
				res[i] = root
				break
			}
			if res[i] == "" {
				res[i] = ss[i]
			}
		}
	}
	return strings.Join(res, " ")

}
