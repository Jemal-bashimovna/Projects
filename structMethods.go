package main

import (
	"fmt"
	"time"
)

// example 1
type Userlist struct {
	login, email, birthday string
	phone, age             int
	score                  []int
}

func (u Userlist) getPhone() int {
	return u.phone
}

func (u *Userlist) setPhone(number1 int) {
	u.phone = number1
}

func (u Userlist) isElder() bool {
	a := u.age
	isTrue := false

	if a >= 18 {
		isTrue = true
	} else if a < 18 {
		isTrue = false
	}
	return isTrue
}

func (u Userlist) getHighScore() int {
	highScore := 0
	for _, score := range u.score {
		if highScore < score {
			highScore = score
		}
	}
	return highScore
}

// example 2

type Nperson struct {
	Nname string
	Nage  int
}

func (p Nperson) printName() {
	fmt.Println(p.Nname)
}

type Builder struct {
	Nperson
	Nname string
	WorkBuilders
}

func (b Builder) printName() {
	fmt.Println(b.Nname)
}

type WorkBuilders struct {
	Nname string
	Nage  int
}

func main() {
	user := Userlist{"mark0990", "mark@mail.com", time.Now().String(), 20101, 19, []int{12, 23, 34, 45}}
	fmt.Println(user)
	fmt.Println(user.getPhone())
	user.setPhone(998899)
	fmt.Println(user.getPhone())

	// checking age

	if user.isElder() {
		fmt.Println("Ok")
	} else {
		fmt.Println("NO")
	}

	fmt.Println(user.getHighScore())

	// Example 2

	person1 := Nperson{Nname: "Person1", Nage: 35}
	person1.printName() // wyzow metoda

	//builder1 := Builder{Nperson{"PersonName", 40}}
	//builder1 := Builder{Nperson{"PersonName", 40}, "BuilderName"} // shadowing example
	builder1 := Builder{
		Nperson{"PersonName", 45},
		"BuilderName",
		WorkBuilders{"Driver", 5}} // colliding example

	// shorthands
	fmt.Println(builder1.Nperson.Nage)
	//	fmt.Println(builder1.Nage)
	fmt.Println(builder1)
	//fmt.Printf("Type: %T, value: %#v\n", builder1, builder1)

	// Shadowing
	fmt.Println("\nShadowing example:\n")
	fmt.Println(builder1.Nname)
	builder1.printName()
	builder1.printName()
	builder1.Nperson.printName()

	// colliding - stalkiwayushiysya
	fmt.Println("\nColliding example:\n")
	fmt.Println(builder1.Nperson.Nname)
	fmt.Println(builder1.Nname)
	fmt.Println(builder1.WorkBuilders.Nname)
	builder1.Nperson.printName()
	builder1.printName()

}
