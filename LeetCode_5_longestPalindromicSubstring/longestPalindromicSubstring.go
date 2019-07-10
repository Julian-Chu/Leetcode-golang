package LeetCode_5_longestPalindromicSubstring

import (
	"fmt"
	"math"
)

func longestPalindrome(s string) string {
	arr := []rune(s)
	length := 0
	subs := ""

	for start := 0; start < len(arr); start++ {
		subarr := arr[start:]
		lenSubarr := len(subarr)
		for i := 0; i < lenSubarr; i++ {
			newSubArr := subarr[:lenSubarr-i]
			newLenSubArr := len(newSubArr)
			mid := 0
			var left []rune
			var right []rune
			isEven := newLenSubArr%2 == 0
			if isEven {
				mid = newLenSubArr/2 - 1
				left = newSubArr[:mid+1]
				right = newSubArr[mid+1:]

			} else {
				mid = int(math.Floor(float64(newLenSubArr / 2)))
				left = newSubArr[:mid+1]
				right = newSubArr[mid:]
			}
			fmt.Println(left)
			fmt.Println(right)
			fmt.Printf("mid: %v\n", mid)
			if !arrIsEqual(mid, left, right) {
				continue
			}

			if newLenSubArr > length {
				length = newLenSubArr
				fmt.Printf("sub arr len: %v\n", length)
				if isEven {
					subs = string(append(left, right...))
				} else {
					subs = string(append(left[:len(left)-1], right...))
				}
			}

		}
	}
	return subs
}

func arrIsEqual(mid int, left []rune, right []rune) bool {
	if len(left) == 0 || len(right) == 0 {
		return false
	}

	for j := 0; j <= mid; j++ {
		if left[j] != right[len(right)-j-1] {
			return false
		}
	}

	if mid == 0 {
		if left[0] != right[0] {
			return false
		}
	}
	return true
}
