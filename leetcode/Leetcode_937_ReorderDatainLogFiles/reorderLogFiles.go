package Leetcode_937_ReorderDatainLogFiles

import (
	"sort"
	"strings"
)

func reorderLogFiles(logs []string) []string {
	type logStruct struct {
		ident   string
		content string
		typ     int
		idx     int
	}

	const (
		letter = iota
		digit  = iota
	)

	parse := func(idx int, log string) logStruct {
		arr := strings.SplitN(log, " ", 2)

		l := logStruct{
			arr[0],
			arr[1],
			letter,
			idx,
		}

		if arr[1][0] >= '0' && arr[1][0] <= '9' {
			l.typ = digit
		}
		return l
	}

	logStructs := make([]logStruct, 0, len(logs))
	for i := range logs {
		logStructs = append(logStructs, parse(i, logs[i]))
	}

	sort.SliceStable(logStructs, func(i, j int) bool {
		first := logStructs[i]
		second := logStructs[j]
		if first.typ != second.typ {
			return first.typ < second.typ
		}

		if first.typ == digit {
			return false
		}

		if first.content == second.content {
			return first.ident < second.ident
		}

		return first.content < second.content
	})

	res := make([]string, 0, len(logs))
	for _, log := range logStructs {
		res = append(res, logs[log.idx])
	}

	return res
}
