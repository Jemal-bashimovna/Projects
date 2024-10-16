package main

import (
	"fmt"
	"time"
)

type ourString string
type ourInt int

type users struct {
	name     string
	email    string
	password int
}

func main() {
	var name ourString
	var age ourInt
	fmt.Printf("%T %#v \n", name, name)
	fmt.Printf("%T %#v \n", age, age)

	name = "Kate"
	age = 25
	fmt.Printf("%T %#v \n", name, name)
	fmt.Printf("%T %#v \n", age, age)

	// create default

	var John users
	fmt.Printf("%T %#v \n", John, John)
	//John = users{}
	//fmt.Printf("%T %#v \n", John, John)

	John.name = "John"
	John.email = "gggg@mail.com"
	John.password = 12345
	fmt.Printf("%T %#v \n", John, John)
	fmt.Println(John)

	// 1- atlaryny gorkezip bahalandyrmak
	user2 := users{
		name:     "Brad",
		email:    "ggjh@mail.com",
		password: 345,
	}

	fmt.Println(user2)

	// 2 - atlaryny gorkezman bahalandyrmak

	user3 := users{"Kate", "kate@mail.com", 987654}
	fmt.Println(user3)

	//доступ к полю через указатель
	pJohn := &John

	fmt.Println((*pJohn).password) // wariant 1
	fmt.Println(pJohn.password)    // wariant 2

	// create pointer to struct
	pUser4 := &users{"user4", "user@mail.com", 88888}
	fmt.Println(pUser4)

	anonymStrtuct := struct {
		login, password, birthday string
		phone                     int
	}{

		login:    "userlogin",
		password: "000tt",
		// birthday: fmt.Sprintf("%s\n", time.Now()),
		birthday: time.Now().String(),
		phone:    0000,
	}

	fmt.Println(anonymStrtuct)
}
