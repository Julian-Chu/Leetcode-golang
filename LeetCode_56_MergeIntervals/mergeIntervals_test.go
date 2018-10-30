package merge_intervals

import (
	"errors"
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
		}

		compareObj := res[len(res)-1]
		fmt.Println(compareObj)
		mergedInv, err := mergeInterval(compareObj, interval)
		if err == nil {
			res[len(res)-1] = mergedInv
			continue
		}
		res = append(res, interval)
	}
	return res
}

func mergeInterval(src Interval, dst Interval) (Interval, error) {
	if src.End < dst.Start {
		return Interval{}, errors.New("no merge")
	}
	var newStart, newEnd int
	newStart = dst.Start
	if src.Start <= dst.Start {
		newStart = src.Start
	}
	newEnd = dst.End
	if src.End >= dst.End {
		newEnd = src.End
	}
	return Interval{newStart, newEnd}, nil

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
	fmt.Println(res)
	if !reflect.DeepEqual(res, []Interval{Interval{1, 3}, Interval{4, 5}}) {
		t.Error("Failed")
	}
}
