package syntex

import (
	"fmt"
	fmt2 "go-sample/builtin/fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

func SampleInterface() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex2{3, 4}

	a = f  // a MyFloat implements Abser
	a = &v // a *Vertex implements Abser

	// In the following line, v is a Vertex (not *Vertex)
	// and does NOT implement Abser.
	a = v

	fmt.Println(a.Abs())
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex2 struct {
	X, Y float64
}

func (v Vertex2) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

/////////////////////////

type I interface {
	M()
}

type T struct {
	S string
}

// This method means type T implements the interface I,
// but we don't need to explicitly declare that it does so.
func (t T) M() {
	fmt.Println(t.S)
}

func SampleInterface2() {
	var i I = T{"hello"}
	i.M()
}

///////////////////////////
/**
if it is empty interface, you can set all value on the interface.
it is similar with Object in Java
*/
func SampleInterface3() {
	var i interface{}
	fmt2.Describe(i)
	i = 42
	fmt2.Describe(i)

	i = "hello"
	fmt2.Describe(i)
}
