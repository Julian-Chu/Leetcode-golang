package string

import "strings"

func trimSpace() {
	_ = strings.TrimSpace(" a b c d e ") // "a b c d e"
	_ = strings.Trim(" a b c d e ", " ") // "a b c d e"
}

func split() {
	_ = strings.Split("a b c d e", " ") // [a b c d e]
}
