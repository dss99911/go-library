package fmt

import "fmt"

func FmtSampleError() {
	e := fmt.Errorf("not found: %s", "http://www.naver.com")
	fmt.Println(e)
}
