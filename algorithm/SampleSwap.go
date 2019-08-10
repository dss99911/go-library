package algorithm

import "fmt"

func SampleSwap() {
	x, y := 0, 1
	x, y = y, x
	fmt.Println(x, y)
}
