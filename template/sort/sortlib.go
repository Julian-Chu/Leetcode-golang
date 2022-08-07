package sort

import "sort"

func sortInts(nums []int) {
	sort.Ints(nums)
	_ = sort.IntsAreSorted(nums)
}

func sortString(strs []string) {
	sort.Strings(strs)
}

func reverse(strs []string) {
	sort.Sort(sort.Reverse(sort.StringSlice(strs)))
	ints := []int{1, 2, 3, 4}
	sort.Sort(sort.Reverse(sort.IntSlice(ints)))
}
