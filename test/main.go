package main

import (
	"fmt"
	"time"
)

func init() {
}
func main() {
	for i := 0; i < 1000000; i++ {
		go func(input int) {
			t := 1000000 - input
			time.Sleep(time.Duration(t) * time.Millisecond)
			fmt.Printf("ABCDEFGIJKLMNOPQRSTZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZZ\n")
		}(i)
	}
	var str string
	fmt.Println("-------")
	fmt.Scanln(&str)
}
