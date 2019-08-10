package syntex

import "fmt"

/**
if use pointer, you can change field
*/
func (u *User) setEmail(email string) {
	u.Email = email
}

/**
if use without pointer, make new struct, so, even if change field, it is not reflected on origin
*/
func (u User) GetEmail3() string {

	return u.Email
}

func SampleStructExtenstionMethod() {
	user := User{}
	email2 := user.GetEmail3()
	fmt.Println(email2)
}

type Test struct {
	S string
}

/**
nil is possible for pointer variable
*/
func (t *Test) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}
