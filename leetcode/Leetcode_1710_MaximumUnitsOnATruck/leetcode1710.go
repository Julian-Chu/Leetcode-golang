package Leetcode_1710_MaximumUnitsOnATruck

import "container/heap"

type Heap [][]int

func (h *Heap) Len() int {
	return len(*h)
}

func (h *Heap) Less(i, j int) bool {
	return (*h)[i][1] > (*h)[j][1]
}

func (h *Heap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *Heap) Push(x interface{}) {
	*h = append(*h, x.([]int))
}

func (h *Heap) Pop() interface{} {
	r := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return r
}

func maximumUnits(boxTypes [][]int, truckSize int) int {
	h := Heap(boxTypes)
	heap.Init(&h)

	r := 0
	for truckSize > 0 && h.Len() > 0 {
		tmp := heap.Pop(&h).([]int)
		if tmp[0] <= truckSize {
			r += tmp[0] * tmp[1]
			truckSize -= tmp[0]
			continue
		}

		r += truckSize * tmp[1]
		break
	}
	return r
}

//func maximumUnits(boxTypes [][]int, truckSize int) int {
//	buckets := make([]int, 1001)
//
//	for _, box := range boxTypes {
//		buckets[box[1]] += box[0]
//	}
//
//	units := 0
//
//	for i := 1000; i > 0; i-- {
//		if truckSize >= buckets[i] {
//			units += buckets[i] * i
//			truckSize -= buckets[i]
//		} else {
//			units += truckSize * i
//			return units
//		}
//	}
//	return units
//}

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
