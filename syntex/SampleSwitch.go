package syntex

import "fmt"

func SampleSwitch(a int) {
	switch a {
	case 1:
		fmt.Println("a>1")
		fallthrough //doesn't break but go throw case 0
	case 0:
		fmt.Println("a == 0")
	default:
		fmt.Println("a < 0")
	}
}

func SampleSwitch2() {
	switch a := 1; a {
	case 1:
		fmt.Println("a>1")
	case 0:
		fmt.Println("a == 0")
	default:
		fmt.Println("a < 0")
	}
}

func SampleSwitchWithNoCondition(a int) {
	switch {
	case a > 1:
		fmt.Println("a>1")
	case a == 0:
		fmt.Println("a == 0")
	default:
		fmt.Println("a < 0")
	}
}
