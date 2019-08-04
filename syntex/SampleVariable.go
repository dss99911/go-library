package syntex

var x = 2  // declare globar variable x
var a bool // defualt false
var b int  // default 0
var i2, j2 = 1, 2

const Pi = 3.14

func SampleVariable() {
	var y int

	//declare implicit local variable. it is not possible for global variables.
	x := 2
	var c, python, java = true, false, "no!"
	const World = "世界"

	println(x, y, a, b, c, python, java)

	i := 42           // int
	f := 3.142        // float64
	g := 0.867 + 0.5i // complex128
	println(i, f, g)
}

type SampleType struct {
	bool

	string

	int
	int8
	int16
	int32
	int64

	uint
	uint8
	uint16
	uint32
	uint64
	uintptr

	byte // alias for uint8

	rune // alias for int32
	// represents a Unicode code point

	float32
	float64

	complex64
	complex128
}
