package main

import (
	"fmt"
)

type Person struct {
	Name, Phone, Email, Address string
}

func validatePhoneNumber(phone string) bool {

	if !(len(phone) == 12 && phone[0:4] == "+993") {
		return false
	}
	for i := 1; i < 12; i++ {
		if !(phone[i] >= '0' && phone[i] <= '9') {
			return false
		}
	}
	return true
}

func (p *Person) SetPhone(newPhone string) {

	if validatePhoneNumber(newPhone) {
		p.Phone = newPhone
	}

}

func validateEmail(email string) bool {

	if !(len(email) <= 500 && email[0] != '@' && email[len(email)-1] != '@') {
		return false
	}
	num := 0
	//validCount := 1
	for i := 1; i < len(email)-1; i++ {
		if email[i] == '@' {
			num++
		}
	}
	return num == 1
}

//func isEmailValid(e string) bool {
//	emailRegex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
//	return emailRegex.MatchString(e)
//}

func (p *Person) SetEmail(newEmail string) {
	if validateEmail(newEmail) {
		p.Email = newEmail
	}
}

func main() {

	person := Person{"John", "+99361334400", "exampl@g", "London"}

	validPhoneNumber := validatePhoneNumber(person.Phone)
	fmt.Println(validPhoneNumber)

	if validPhoneNumber != false {
		person.Phone = "+99361004455"
	}
	fmt.Println(person)

	person.SetPhone("+993aaaaaaaa")
	fmt.Println(person)

	validEmail := validateEmail(person.Email)
	fmt.Println(validEmail)

	person.SetEmail("example@email.com")
	fmt.Println(person)

}
