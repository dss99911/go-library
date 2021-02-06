package syntex

import "fmt"

func SampleClosure() {
	f := fibonacci4()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

func fibonacci4() func() int {
	a1,a2 := 0, 1
	return func() int {
		result := a1
		a1, a2 = a2, a1+a2
		return result
	}
}
