package main

import (
	"fmt"
	"math"
	"testing"
)

func poorPigs(buckets, minutesToDie, minutesToTest int) int {
	times := (int)(minutesToTest / minutesToDie)
	pigs := 0
	count := 0
	for count < buckets {
		if buckets == 1 {
			return 0
		}
		pigs++
		count = int(math.Pow(float64(pigs+1), float64(times)))
	}

	return pigs
}

func Test_case1(t *testing.T) {
	res := poorPigs(1000, 15, 60)
	if res != 5 {
		fmt.Println("res:", res)
		t.Error("Failed", res)
	}
}

func Test_case2(t *testing.T) {
	res := poorPigs(1, 15, 60)
	if res != 0 {
		fmt.Println("res:", res)
		t.Error("Failed", res)
	}
}

func Test_case3(t *testing.T) {
	res := poorPigs(1000, 12, 60)
	if res != 4 {
		fmt.Println("res:", res)
		t.Error("Failed", res)
	}
}
