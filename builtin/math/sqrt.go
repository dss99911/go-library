package math

import (
	"fmt"
	"math"
	"math/cmplx"
)

func SampleMath() {
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(float64(5*5+10*10)))
	fmt.Printf("Now you have %g problems.\n", cmplx.Sqrt(-5+12i))
	fmt.Println(math.Pi)

}
