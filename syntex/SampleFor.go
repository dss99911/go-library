package syntex

import "fmt"

func SampleFor() {
	Sum()
	sum2()
	sum3()
	infinite()
}

func Sum() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
		if i > 6 {
			break
		}
	}
	fmt.Println(sum)
}

func sum2() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

func sum3() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

func infinite() {
	for {
		fmt.Println("d")
	}
}
