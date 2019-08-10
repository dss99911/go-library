package syntex

import "fmt"

func SampleGlobalFunction() {
	var s []int
	fmt.Println(s, len(s), cap(s))

	c := make(chan int, 10)
	c <- 1
}
