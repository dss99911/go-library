package syntex

import (
	"fmt"
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
