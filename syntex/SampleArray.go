package syntex

import (
	"fmt"
	"strings"
)

func SampleArray() {
	var aa [2]string
	aa[0] = "Hello"
	aa[1] = "World"
	fmt.Println(aa[0], aa[1])

	a := [5]int{1, 2, 3, 4, 5}
	fmt.Println(a)

	//append
	myslice := []int{1, 2, 3, 4, 5} //[]int{}
	myslice = append(myslice, 6, 7, 8)
	myslice = append(myslice, 9)
	fmt.Println(myslice)

	//slice
	s := []int{1, 2, 3, 4, 5} //[]int{}
	s1 := s[1:3]
	fmt.Println(s1[:cap(s1)])
	fmt.Println(gettwo(s, 1, 3))

	printSlice(s1)
	nilArray()
	useMake()

	multiSlice()
}

func multiSlice() {
	// Create a tic-tac-toe board.
	board := [][]string{
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	}
	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"
	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

func useMake() {
	a := make([]int, 5)
	printSlice(a)

	b := make([]int, 0, 5)
	printSlice(b)

	c := b[:2]
	printSlice(c)

	d := c[2:5]
	printSlice(d)
}

func nilArray() {
	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil!")
	}
}

func gettwo(s []int, fi, si int) []int { // if s has 100 elements, then the return value will have access to 100-fi elements
	s2 := make([]int, si-fi)
	copy(s2, s[fi:si])
	return s2
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
