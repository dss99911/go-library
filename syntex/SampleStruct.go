package syntex

import "fmt"

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

type fieldNameWithType struct {
	int
	string
	person //if doesn't mention field name, it is similar to inherit the type. so, fieldNameWithType can use person's method directly.
	User
}

func SampleStruct() {
	jason := person{
		name:    "Json S.",
		age:     38,
		address: "Germany",
	}

	//without field name
	test := person{"sdaf", 23, "sdf"}

	//set default value
	test2 := person{}

	fmt.Println(jason.name, test.name, test2.name)

	user := User{
		Id:    "dsf",
		Email: "asdf",
	}
	fmt.Println(user.GetEmail())

	user2 := NewUser("kim", "a@gmail.com")
	fmt.Println(user2.GetEmail3())

	withType := fieldNameWithType{
		int:    0,
		string: "",
		person: person{
			name:    "sdf",
			age:     0,
			address: "sdaf",
		},
		User: User{
			Id:    "asdf",
			Email: "asdf",
		},
	}

	print(withType.GetEmail3())
}

func NewUser(id, email string) *User {
	return &User{
		Id:    id,
		Email: email,
	}
}

/**
there is no way to set default value.
only this way. we can set some default value
*/
func NewDefaultUser() *User {
	return &User{
		Id:    "d",
		Email: "test@gmail.com",
	}
}

func (u *User) GetID() string {
	return u.Id
}

func (u User) GetEmail() string {
	return u.Email
}
