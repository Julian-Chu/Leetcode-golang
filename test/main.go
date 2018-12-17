package main

import (
	"bytes"
	"fmt"
	"strconv"
	"time"
)

func main() {
	// var buffer bytes.Buffer
	// for i := 0; i < 10; i++ {
	// 	go func(input int) {
	// 		fmt.Println(input)
	// 		buffer.WriteString(strconv.Itoa(input))
	// 		time.Sleep(500 * time.Millisecond)
	// 		fmt.Println(buffer.String())
	// 	}(i)
	// }

	// for i := 0; i < 10; i++ {

	// 	var buffer bytes.Buffer
	// 	go func() {
	// 		// wg.Add(1)
	// 		// time.Sleep(500 * time.Microsecond)
	// 		fmt.Println(buffer.String())
	// 		// wg.Done()
	// 	}()
	// 	for j := 0; j < 30; j++ {
	// 		buffer.WriteString(strconv.Itoa(i))
	// 	}
	// }

	// log.Printf("%s", "1.1.1.002")
	// sigTerm := syscall.Signal(15)

	// log.Println(reflect.TypeOf(sigTerm))
	// fmt.Println(4)
	// fmt.Println(5)

	fmt.Println("Hello, playground")
	c := make(chan int, 3)
	c <- 1
	c <- 2
	c <- 3

	go func(c chan int) {
		for v := range c {
			fmt.Println("value:", v)
		}
		c <- 4

	}(c)
	//var input string
	//fmt.Scanf("\n",&input)
	time.Sleep(2 * time.Second)

	v, ok := <-c
	fmt.Println("value:", v, "ok:", ok)
	fmt.Println("before")

	close(c)
	// var str string
	// fmt.Scanln(&str)
}

func Count() (res string) {
	var buffer bytes.Buffer
	for i := 0; i < 1000000; i++ {
		buffer.WriteString(strconv.Itoa(i))
	}
	return buffer.String()
}
