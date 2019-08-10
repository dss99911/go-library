package algorithm

import "fmt"

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main() {
	x, y := 0, 1
	x, y = y, x
	fmt.Println(x, y)

	c := make(chan int, 10)
	fmt.Println(cap(c))
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}