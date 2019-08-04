package syntex

import "fmt"

func SampleIf(a int) {
	if a > 1 {
		fmt.Println("a>1")
	} else if a == 0 {
		fmt.Println("a == 0")
	} else {
		fmt.Println("a < 0")
	}

}
