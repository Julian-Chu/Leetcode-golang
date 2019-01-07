package merge_intervals

import (
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

// {2, 3},  {8, 9}, {1, 10}
// {1, 10}
func merge(intervals []Interval) []Interval {
	res := sort(intervals)
	last := len(res) - 1
	for i := last; i > 0; i-- {
		if res[i-1].End < res[i].Start || res[i-1].Start > res[i].End {
			continue
		}

		var newStart, newEnd int
		newStart = res[i].Start
		if res[i-1].Start <= res[i].Start {
			newStart = res[i-1].Start
		}
		newEnd = res[i].End
		if res[i-1].End >= res[i].End {
			newEnd = res[i-1].End
		}
		res[i-1] = Interval{newStart, newEnd}

		res = append(res[:i], res[(i+1):]...)

	}
	i := 1
	for i < len(res) {
		if res[i-1].End < res[i].Start || res[i-1].Start > res[i].End {
			i++
			continue
		}

		var newStart, newEnd int
		newStart = res[i].Start
		if res[i-1].Start <= res[i].Start {
			newStart = res[i-1].Start
		}
		newEnd = res[i].End
		if res[i-1].End >= res[i].End {
			newEnd = res[i-1].End
		}
		res[i-1] = Interval{newStart, newEnd}

		res = append(res[:i], res[(i+1):]...)
	}
	return res
}

func sort(intervals []Interval) []Interval {
	if len(intervals) <= 1 {
		return intervals
	}
	start := make([]Interval, 0)
	end := make([]Interval, 0)
	m := intervals[int(len(intervals)/2)]

	for _, item := range intervals {
		switch {
		case item.Start < m.Start:
			start = append(start, item)
		case item.Start == m.Start:
			if item.End < m.End {
				start = append(start, item)
			} else if item.End > m.End {
				end = append(end, item)
			}
		case item.Start > m.Start:
			end = append(end, item)
		}
	}

	start, end = sort(start), sort(end)
	start = append(start, m)
	start = append(start, end...)
	return start

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
	if !reflect.DeepEqual(res, []Interval{{1, 10}}) {
		t.Error("Failed")
	}
}

// [[4,5],[2,4],[4,6],[3,4],[0,0],[1,1],[3,5],[2,2]]
func Test_case8(t *testing.T) {
	input := []Interval{{4, 5}, {2, 4}, {4, 6}, {3, 4}, {0, 0}, {1, 1}, {3, 5}, {2, 2}}
	res := merge(input)
	if !reflect.DeepEqual(res, []Interval{{0, 0}, {1, 1}, {2, 6}}) {
		t.Error("Failed")
	}
}

// [[2,3],[5,5],[2,2],[3,4],[3,4]]
// [[2,4],[5,5]]
func Test_case9(t *testing.T) {
	input := []Interval{{2, 3}, {5, 5}, {2, 2}, {3, 4}, {3, 4}}
	res := merge(input)
	if !reflect.DeepEqual(res, []Interval{{2, 4}, {5, 5}}) {
		t.Error("Failed")
	}
}

func Test_case10(t *testing.T) {
	input := []Interval{{2, 3}, {8, 9}, {1, 10}}
	res := merge(input)
	if !reflect.DeepEqual(res, []Interval{{1, 10}}) {
		t.Error("Failed")
	}
}
