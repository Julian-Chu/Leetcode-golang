package Leetcode_937_ReorderDatainLogFiles

import (
	"sort"
	"strings"
)

func reorderLogFiles(logs []string) []string {
	if len(logs) < 2 {
		return logs
	}

	var letterLogs []string
	var digitLogs []string
	for _, log := range logs {
		i := strings.Index(log, " ")
		if log[i+1] >= '0' && log[i+1] <= '9' {
			digitLogs = append(digitLogs, log)
		} else {
			letterLogs = append(letterLogs, log)
		}
	}

	sort.Slice(letterLogs, func(i, j int) bool {
		iIdx := strings.Index(letterLogs[i], " ")
		jIdx := strings.Index(letterLogs[j], " ")

		iLog := letterLogs[i][iIdx+1:]
		jLog := letterLogs[j][jIdx+1:]
		if iLog == jLog {
			return letterLogs[i] < letterLogs[j]
		}
		return iLog < jLog
	})

	letterLogs = append(letterLogs, digitLogs...)
	return letterLogs
}

//func reorderLogFiles(logs []string) []string {
//	sort.SliceStable(logs, func(i, j int) bool {
//		s1 := strings.SplitN(logs[i], " ", 2)
//		s2 := strings.SplitN(logs[j], " ", 2)
//		f1, f2 := "0"+s1[1], "0"+s2[1]
//		if unicode.IsNumber(rune(f1[1])) {
//			f1 = "1"
//		}
//		if unicode.IsNumber(rune(f2[1])) {
//			f2 = "1"
//		}
//		if f1 == f2 {
//			if f1 == "1" {
//				return false
//			}
//
//			return s1[0] < s2[0]
//		}
//		return f1 < f2
//	})
//	return logs
//}

//func reorderLogFiles(logs []string) []string {
//	type logStruct struct {
//		ident   string
//		content string
//		typ     int
//		idx     int
//	}
//
//	const (
//		letter = iota
//		digit  = iota
//	)
//
//	parse := func(idx int, log string) logStruct {
//		arr := strings.SplitN(log, " ", 2)
//
//		l := logStruct{
//			arr[0],
//			arr[1],
//			letter,
//			idx,
//		}
//
//		if arr[1][0] >= '0' && arr[1][0] <= '9' {
//			l.typ = digit
//		}
//		return l
//	}
//
//	logStructs := make([]logStruct, 0, len(logs))
//	for i := range logs {
//		logStructs = append(logStructs, parse(i, logs[i]))
//	}
//
//	sort.SliceStable(logStructs, func(i, j int) bool {
//		first := logStructs[i]
//		second := logStructs[j]
//		if first.typ != second.typ {
//			return first.typ < second.typ
//		}
//
//		if first.typ == digit {
//			return false
//		}
//
//		if first.content == second.content {
//			return first.ident < second.ident
//		}
//
//		return first.content < second.content
//	})
//
//	res := make([]string, 0, len(logs))
//	for _, log := range logStructs {
//		res = append(res, logs[log.idx])
//	}
//
//	return res
//}
