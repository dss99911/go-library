package syntex

import "fmt"

func sum10(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func SampleChannel() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum10(s[:len(s)/2], c)
	go sum10(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)
}

/**
if set buffer, you can set value in advance
if exceed the buffer, error will occur
*/
func SampleBufferedChannel() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

/**
Note: Only the sender should close a channel, never the receiver. Sending on a closed channel will cause a panic.

Another note: Channels aren't like files; you don't usually need to close them. Closing is only necessary when the receiver must be told there are no more values coming, such as to terminate a range loop.
*/
func SampleCloseAndRangeChannel() {
	x, y := 0, 1
	x, y = y, x
	fmt.Println(x, y)

	c := make(chan int, 20)
	fmt.Println(cap(c))
	go fibonacci2(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}

func fibonacci2(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}
