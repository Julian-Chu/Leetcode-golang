package merge_intervals

import (
	"reflect"
	"testing"
)

func Test_case1(t *testing.T) {
	input := []Interval{{1, 4}, {4, 5}}
	res := merge(input)
	if !reflect.DeepEqual(res, []Interval{{1, 5}}) {
		t.Error("Failed")
	}
}

func Test_case2(t *testing.T) {
	input := []Interval{{1, 3}, {4, 5}}
	res := merge(input)
	if !reflect.DeepEqual(res, []Interval{{1, 3}, {4, 5}}) {
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
	input := []Interval{{1, 4}, {0, 0}}
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
