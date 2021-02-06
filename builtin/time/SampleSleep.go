package time

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		//duration unit is nanosecond
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)

		time.Sleep(100 * time.Minute)
	}
}
