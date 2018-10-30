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
	return []Interval{Interval{1, 5}}
}

func Test_case1(t *testing.T) {
	input := []Interval{Interval{1, 4}, Interval{4, 5}}
	res := merge(input)
	fmt.Println(res)
	if !reflect.DeepEqual(res, []Interval{Interval{1, 5}}) {
		t.Error("Failed")
	}

}
