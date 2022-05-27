package Leetcode_1710_MaximumUnitsOnATruck

func maximumUnits(boxTypes [][]int, truckSize int) int {
	buckets := make([]int, 1001)

	for _, box := range boxTypes {
		buckets[box[1]] += box[0]
	}

	units := 0

	for i := 1000; i > 0; i-- {
		if truckSize >= buckets[i] {
			units += buckets[i] * i
			truckSize -= buckets[i]
		} else {
			units += truckSize * i
			return units
		}
	}
	return units
}

//func maximumUnits(boxTypes [][]int, truckSize int) int {
//	sort.Slice(boxTypes, func(i, j int) bool {
//		if boxTypes[i][1] > boxTypes[j][1] {
//			return true
//		}
//		return false
//	})
//
//	var cnt = 0
//
//	for _, boxType := range boxTypes {
//		if truckSize <= 0 {
//			break
//		}
//		totalNum := boxType[0]
//		unitsPerBox := boxType[1]
//		num := totalNum
//		if truckSize < num {
//			num = truckSize
//		}
//		cnt += num * unitsPerBox
//		truckSize -= num
//	}
//	return cnt
//}
