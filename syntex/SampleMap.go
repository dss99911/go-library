package syntex

import "fmt"

func SampleMap() {
	// key ==> value (key value pair)
	/*
		x := make(map[string]int)
		x["first"] = 1
		x["second"] = 2
	*/
	SampleMap1()

	x := SampleMap2()

	fmt.Println(x["first"])
	if v, ok := x["second"]; ok {
		fmt.Println(v)
	}
	fmt.Println(x)
	delete(x, "first")
	fmt.Println(x)

	//no need Vertex struct name
	var m = map[string]Vertex{
		"Bell Labs": {40, -74},
		"Google":    {37, -122},
	}
	fmt.Println(m)

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}

func SampleMap1() {
	m := make(map[string]int)
	m["Bell Labs"] = 1
	fmt.Println(m["Bell Labs"])
}

func SampleMap2() map[string]int {
	x := map[string]int{
		"first":  1,
		"second": 2,
	}
	return x
}
