package main

import (
	"fmt"
	"strings"
)

func main() {

	message1 := `Hello`
	fmt.Println(message1)
	message := "Hello world"
	fmt.Println(message)
	// access first character
	fmt.Printf("%c", message[0])
	// access second character
	fmt.Printf("%c\n", message[1])

	// use len() function to count length
	stringLen := len(message)
	fmt.Println("Length of string : ", stringLen)

	result := message1 + " " + message
	fmt.Println(result)

	//Методы работы со строками в Golang
	/*

		Функции	Описания
		Compare()	сравнивает две строки
		Contains()	проверяет, присутствует ли подстрока внутри строки
		Replaces()	заменяет подстроку другой подстрокой
		ToLower()	преобразует строку в нижний регистр
		ToUpper()	преобразует строку в верхний регистр
		Split()	разбивает строку на несколько подстрок

	*/

	string1 := "golang"
	string2 := "\"Bashimowna\"Jemal"
	string3 := "GO PROGRAMMING"

	fmt.Println(strings.Compare(string1, string2))
	fmt.Println(strings.Compare(string2, string3))
	fmt.Println(strings.Compare(string1, string3))

	string4 := "Go Programming"
	// check if string1 (Go) is present in srting2 (Go Programming)
	text := strings.Contains(string4, string1)
	fmt.Println(text)
	//
	fmt.Println(string2)
	replacedString2 := strings.Replace(string2, "a", "X", 10)
	fmt.Println(replacedString2)

	fmt.Println(strings.ToUpper(string1))
	fmt.Println(strings.ToLower(string3))
	fmt.Println(strings.Split(string2, " / "))
}
