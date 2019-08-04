package syntex

import "fmt"

func SampleStruct() {
	jason := person{
		name:    "Json S.",
		age:     38,
		address: "Germany",
	}
	fmt.Println(jason.name)

	user := User{
		Id:    "dsf",
		Email: "asdf",
	}
	fmt.Println(user.GetEmail())

	user2 := NewUser("kim", "a@gmail.com")
	fmt.Println(user2.GetEmail2())
}

/**
private class if start with lower character
*/
type person struct {
	name    string
	age     int
	address string
}

type User struct {
	Id    string
	Email string
}

func NewUser(id, email string) *User {
	return &User{
		Id:    id,
		Email: email,
	}
}

func (u *User) GetID() string {
	return u.Id
}

func (u *User) GetEmail() string {
	return u.Email
}
