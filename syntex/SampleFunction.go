package syntex

import (
	"fmt"
	"math"
)

func SampleFuction() {
	fmt.Println(add(42, 13))
	b, i := CompareNumbers(1, 2)
	fmt.Printf("%t %d", b, i)
}

func add(x int, y int) int {
	return x + y
}

func CompareNumbers(i1, i2 int) (bool, int) {
	if i1 > i2 {
		return false, i1 - i2
	} else if i2 > i1 {
		return false, i2 - i1
	}

	return true, 0
}

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func SampleFunctionParameter() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}

func SampleFunctionClosure() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func SampleFibonacci() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

func fibonacci() func() int {
	a1, a2 := 0, 1
	return func() int {
		result := a1
		a1 = a2
		a2 = result + a1

		return result
	}
}
