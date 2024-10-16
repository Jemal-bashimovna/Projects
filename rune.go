package main

import (
	"fmt"
	"reflect"
)

func main() {
	// simple rune
	var r rune = 'A'
	fmt.Println(r)

	// rune array
	str := "abcdefg"
	runeArray := []rune(str)
	fmt.Println("The rune array:")
	fmt.Println(runeArray)

	//
	rune1 := 'a'
	rune2 := 't'
	fmt.Printf("Unicode: %U ; type: %s\n", rune1, reflect.TypeOf(rune1))
	fmt.Printf("Unicode: %U ; type: %s\n", rune2, reflect.TypeOf(rune2))

	//
	res := Reverse("hello")
	fmt.Println(res)
	//

	res2 := Reverse2("Hello Guys")
	fmt.Println(res2)
}

func Reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}

func Reverse2(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
