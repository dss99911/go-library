package fmt

import (
	"fmt"
	"io"
	"strings"
)

func SampleFmt() {
	fmt.Println("test")

	//print several variables
	fmt.Println("test", 1, 2, 3)
	fmt.Printf("Type: %T Value: %v\n", 1, 1)
}

func Describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

/**
%q looks change []byte to string
*/
func main() {
	r := strings.NewReader("Hello, Reader!")

	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}
