package syntex

import (
	"fmt"
	"sync"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func SampleGoRoutine() {
	go say("world")
	say("hello")
}


func SampleWaitGoRountineCompleted() {
	sleepSeveral(10)
}

func sleepSeveral(count int) {
	var wg sync.WaitGroup
	defer wg.Wait()

	for i:= 0; i < count; i++ {
		wg.Add(1)
		go func() {
			sleep()
			wg.Done()
		}()
		wait()
	}
}

func sleep() {
	time.Sleep(time.Second * 3)
	fmt.Println("wake up")
}

func wait() {
	time.Sleep(time.Second * 1)
	fmt.Println("wait")
}