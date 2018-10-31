package merge_intervals

import (
	"fmt"
	"reflect"
	"testing"
)

/**
 * Definition for an interval.
 * type Interval struct {
 *	   Start int
 *	   End   int
 * }
 */

type Interval struct {
	Start int
	End   int
}

func merge(intervals []Interval) []Interval {
	res := make([]Interval, 0)
	for _, interval := range intervals {
		if len(res) == 0 {
			res = append(res, interval)
			continue
		}
		last := len(res) - 1
		if !(res[last].End < interval.Start || res[last].Start > interval.End) {
			res[last] = mergeInterval(res[last], interval)
			if len(res) > 1 {
				for i := len(res) - 1; i >= 1; i-- {
					if res[i-1].End < res[i].Start || res[i-1].Start > res[i].End {
						continue
					}
					res[i-1] = mergeInterval(res[i-1], res[i])
					res = res[:i]
				}
			}
		} else {
			res = append(res, interval)
			length := len(res)
			if res[length-2].Start > res[length-1].End {
				res[length-2], res[length-1] = res[length-1], res[length-2]
			}
		}
	}
	return res
}

func mergeInterval(src Interval, dst Interval) Interval {
	var newStart, newEnd int
	newStart = dst.Start
	if src.Start <= dst.Start {
		newStart = src.Start
	}
	newEnd = dst.End
	if src.End >= dst.End {
		newEnd = src.End
	}
	return Interval{newStart, newEnd}
}
func Test_case1(t *testing.T) {
	input := []Interval{Interval{1, 4}, Interval{4, 5}}
	res := merge(input)
	if !reflect.DeepEqual(res, []Interval{Interval{1, 5}}) {
		t.Error("Failed")
	}
}

func Test_case2(t *testing.T) {
	input := []Interval{Interval{1, 3}, Interval{4, 5}}
	res := merge(input)
	if !reflect.DeepEqual(res, []Interval{Interval{1, 3}, Interval{4, 5}}) {
		t.Error("Failed")
	}
}

// Input: [[1,3],[2,6],[8,10],[15,18]]
// Output: [[1,6],[8,10],[15,18]]
func Test_case3(t *testing.T) {
	input := []Interval{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	res := merge(input)
	if !reflect.DeepEqual(res, []Interval{{1, 6}, {8, 10}, {15, 18}}) {
		t.Error("Failed")
	}
}
func Test_case4(t *testing.T) {
	input := []Interval{Interval{1, 4}, Interval{0, 0}}
	res := merge(input)
	if !reflect.DeepEqual(res, []Interval{{0, 0}, {1, 4}}) {
		t.Error("Failed")
	}
}

func Test_case5(t *testing.T) {
	input := []Interval{{1, 3}, {4, 5}, {3, 4}}
	res := merge(input)
	if !reflect.DeepEqual(res, []Interval{{1, 5}}) {
		t.Error("Failed")
	}
}

func Test_case6(t *testing.T) {
	input := []Interval{{1, 3}, {4, 5}, {3, 4}, {2, 6}}
	res := merge(input)
	if !reflect.DeepEqual(res, []Interval{{1, 6}}) {
		t.Error("Failed")
	}
}

// [[2,3],[4,5],[6,7],[8,9],[1,10]]
func Test_case7(t *testing.T) {
	input := []Interval{{2, 3}, {4, 5}, {6, 7}, {8, 9}, {1, 10}}
	res := merge(input)
	fmt.Println(res)
	if !reflect.DeepEqual(res, []Interval{{1, 10}}) {
		t.Error("Failed")
	}
}

// [[4,5],[2,4],[4,6],[3,4],[0,0],[1,1],[3,5],[2,2]]
func Test_case8(t *testing.T) {
	input := []Interval{{4, 5}, {2, 4}, {4, 6}, {3, 4}, {0, 0}, {1, 1}, {3, 5}, {2, 2}}
	res := merge(input)
	fmt.Println(res)
	if !reflect.DeepEqual(res, []Interval{{0, 0}, {1, 1}, {2, 6}}) {
		t.Error("Failed")
	}
}
