package Leetcode_1710_MaximumUnitsOnATruck

import (
	"sort"
)

func maximumUnits(boxTypes [][]int, truckSize int) int {
	sort.Slice(boxTypes, func(i, j int) bool {
		if boxTypes[i][1] > boxTypes[j][1] {
			return true
		}
		return false
	})

	var cnt = 0

	for _, boxType := range boxTypes {
		if truckSize <= 0 {
			break
		}
		totalNum := boxType[0]
		unitsPerBox := boxType[1]
		num := totalNum
		if truckSize < num {
			num = truckSize
		}
		cnt += num * unitsPerBox
		truckSize -= num
	}
	return cnt
}
