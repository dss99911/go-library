package syntex

import "fmt"

func sum10(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	fmt.Println("sum10 start")
	c <- sum // send sum to c
	fmt.Println("sum10 end")
}

/**
  - one of sending or receiving should be in the goroutine
  - when receiving, if there is no value in the channel, it waits.
 */
func SampleChannel() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum10(s[:len(s)/2], c)
	go sum10(s[len(s)/2:], c)
	x, y := <-c, <-c//wait until goroutine send by channel
	//x = <- c // if call this, as there is no goroutine working, it crashes
	fmt.Println(x, y, x+y)
}
/**
- when sending, if receiver is not there, then wait.
 */
func SampleChannelNoReceiver() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum10(s[:len(s)/2], c)
	go sum10(s[len(s)/2:], c)
	//result
	// sum10 start
	// sum10 start
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
if you use range, close() should be falled. if not, range is not stopped.
*/
func SampleCloseAndRangeChannel() {
	c := make(chan int)

	fmt.Println("cap", cap(c))
	go fibonacci2(20, c)
	for i := range c {
		v, ok := <-c //check if channel is closed
		fmt.Println(i)
		fmt.Println(v, ok)
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
