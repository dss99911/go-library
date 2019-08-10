package syntex

import (
	"fmt"
	"go-sample/model"
	"io"
	"strings"
)

func SampleString() {
	var test = "apple"
	fmt.Printf("%d", len(test)) // length
	var s = test[1:]
	print(s)

	strings.HasPrefix(test, "app") //Java's startWith
	strings.HasSuffix(test, "le")  //Java's endWith
}

/**
just println. but Person implemented Stringer interface method.
so, string is cutomized
*/
func main() {
	a := model.Person{"Arthur Dent", 42}
	z := model.Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z)
}

func StringReader() {
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
