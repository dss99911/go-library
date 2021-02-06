package algorithm

import "fmt"

/*
https://tour.golang.org/flowcontrol/8

 */
func SampleNetwonMethod() {
	fmt.Println(Sqrt(16))
}

func Sqrt(x float64) float64 {
	z := 1.0
	for i:=0; i < 10; i++ {
		z-= (z*z - x) / (2*z)
	}
	return z
}