package main

import (
	"fmt"
)

type Users struct {
	name string
	age  int
}

type Scores struct {
	User  Users
	Score []int
}

func (u Scores) getHighScore() int {
	maxScore := 0
	for _, val := range u.Score {
		if maxScore < val {
			maxScore = val
		}
	}
	return maxScore
}

func changeStr(str *string) {
	*str = "Hello World"
}

func changeStruct(u *Users) {
	u.name = "John"
	u.age = 15
}

func (u Users) getName() string {
	return u.name
}

func (u *Users) setName(newName string) {
	u.name = newName
}

func (u Users) isElder() bool {

	if u.age >= 18 {
		return true
	}
	return false

	//age := u.age
	//isTrue := false
	//if age >= 18 {
	//	isTrue = true
	//}
	//return isTrue
}

func main() {
	// default value
	var pointer *string
	fmt.Printf("%T %#v \n", pointer, pointer)
	fmt.Printf("%T %#v \n", pointer, pointer)

	// getting not nill pointers

	var a int64 = 7
	fmt.Printf("%T %#v \n", a, a)

	// pointer na ukazatel a

	var pointerA *int64 = &a // & ampersand
	fmt.Printf("%T %#v %#v\n", pointerA, pointerA, *pointerA)

	// create new pointer

	var newP = new(float64)
	fmt.Printf("%T %#v %#v\n", newP, newP, *newP)

	*newP = 4
	fmt.Printf("%T %#v %#v\n", newP, newP, *newP)

	// Pointer usage (example without pointer)
	num := 3

	squareP(&num)
	fmt.Println(num)

	// Pointer usage example 2

	var wallet1 *int
	fmt.Println(hasWallet(wallet1))

	wallet2 := 0
	fmt.Println(hasWallet(&wallet2))

	wallet3 := 6
	fmt.Println(hasWallet(&wallet3))

	x := 5
	//p := &a
	fmt.Println(&x)

	stringToChange := "Hello"
	fmt.Println(stringToChange)

	changeStr(&stringToChange)
	fmt.Println(stringToChange)

	user := Users{"Harry", 30}
	fmt.Println(user)

	changeStruct(&user)
	fmt.Println(user)

	name := user.getName()
	fmt.Println(name)

	fmt.Println(user.getName())

	user.setName("Mark")
	fmt.Println(user)

	if user.isElder() {
		fmt.Println("You can enter")
	} else {
		fmt.Println("You can't enter")
	}

	insertScore := Scores{user, []int{20, 30, 40, 50}}
	fmt.Println("The highest score of user: ", insertScore.getHighScore())

	fmt.Println(insertScore)

}

func squareP(num *int) {
	*num = *num * *num
}

func hasWallet(money *int) bool {
	return money != nil
}
