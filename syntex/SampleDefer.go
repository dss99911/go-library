package syntex

import "fmt"

/**
A defer statement defers the execution of a function until the surrounding function returns.
*/
func SampleDefer() {
	defer fmt.Println("world")

	fmt.Println("hello")
}

/**
Deferred function calls are pushed onto a stack. When a function returns, its deferred calls are executed in last-in-first-out order.
*/
func SampleDeferStack() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
	//counting done 9 8 7 6 5 4 3 2 1 0
}
