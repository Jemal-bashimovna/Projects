package main

import "fmt"

type User struct {
	Id   int64
	Name string
}

func main() {

	maps()

	//// here string - type of key, int - type of key value
	//var person map[string]int = map[string]int{
	//	"Mark": 23,
	//	"Nick": 25,
	//	"Anna": 67,
	//}
	//fmt.Println(person)
	//fmt.Println(person["Anna"])
	//
	//// short syntax
	//
	//students := map[string]int{
	//	"Mark":  23,
	//	"Nick":  25,
	//	"Anna":  67,
	//	"John":  89,
	//	"Peter": 90,
	//}
	//
	//students["Mark"] = 60    // change value of key
	//delete(students, "Nick") // delete key
	//fmt.Println(students)
	//
	//value, truth := students["Mark"] // barlygyny barlamak ( value - value of key, truth - true or false)
	//fmt.Println("value exist ", value, truth)
	//
	//// Program using range with map
	//subjectMarks := map[string]float32{
	//	"math":    70,
	//	"english": 80,
	//	"russian": 96,
	//}
	//
	//fmt.Println(subjectMarks)
	//for subject, mark := range subjectMarks {
	//	fmt.Println(subject, " : ", mark)
	//}
	//
	//for subject := range subjectMarks {
	//	fmt.Println(subject)
	//}
	//
	//for _, mark := range subjectMarks {
	//	fmt.Println(mark)
	//}
	//
	//// create a map
	//squaredNumber := map[int]int{2: 4, 3: 9, 4: 16, 5: 25}
	//
	//// for-range loop to iterate through each key-value of map
	//for n, squared := range squaredNumber {
	//	fmt.Printf("Square of %d is %d\n", n, squared)
	//}
}

func maps() {
	// default value
	var defaultMap map[int]string
	fmt.Printf("defaultMap => Type: %T, value %#v\n", defaultMap, defaultMap)
	fmt.Printf("Length : %d\n\n", len(defaultMap))
	// default value bilen doredilen map ulanmak amatsyz, vlue berip bolanok

	//slice by make
	mapByMake := make(map[int]string)
	fmt.Printf("mapByMake => Type %T, value %#v\n", mapByMake, mapByMake)
	fmt.Printf("Length : %d\n\n", len(mapByMake))

	//slice by make wit capacity
	mapByMakeWithCap := make(map[int]string, 3)
	fmt.Printf("mapByMakeWithCap => Type %T, value %#v\n", mapByMakeWithCap, mapByMakeWithCap)
	fmt.Printf("Length : %d\n\n", len(mapByMakeWithCap))

	//slice by literal
	mapByLiteral := map[string]int{"London": 25, "Moscow": 30}
	fmt.Printf("mapByLiteral => Type %T, value %#v\n", mapByLiteral, mapByLiteral)
	fmt.Printf("Length : %d\n\n", len(mapByLiteral))

	//slice with new
	mapWithNew := *new(map[string]string)
	fmt.Printf("mapWithNew => Type %T, value %#v\n", mapWithNew, mapWithNew)
	fmt.Printf("Length : %d\n\n", len(mapWithNew))

	// insert value
	mapByMake[1] = "First"
	mapByMake[2] = "Second"
	fmt.Printf("mapByMake => Type %T, value %#v\n", mapByMake, mapByMake)
	fmt.Printf("Length : %d\n\n", len(mapByMake))

	// update value
	mapByMake[1] = "Updated First"
	fmt.Printf("mapByMake => Type %T, value %#v\n", mapByMake, mapByMake)
	fmt.Printf("Length : %d\n\n", len(mapByMake))

	// get map value
	fmt.Println(mapByMake[1])

	// get default map value
	fmt.Println(mapByLiteral["Second"])
	fmt.Println(mapByMake[3])

	// checking value = default or value = 0
	val, ok := mapByLiteral["London"]
	fmt.Printf("Value %d isExist: %t \n\n", val, ok)

	// delete value
	delete(mapByLiteral, "London")
	fmt.Printf("Type %T, value %#v\n\n", mapByLiteral, mapByLiteral)

	// map iteration
	mapForIteration := map[string]int{"First": 1, "Second": 2, "Third": 3, "Fourth": 4}

	for key, value := range mapForIteration {
		fmt.Printf("Key => %s, Value => %d\n", key, value)
	} // her gezek durli yzygiderlik alynyar

	// using map as set (filtr unikalnosti)
	// unique values
	students := []User{
		{
			Id:   1,
			Name: "Anna",
		},
		{
			Id:   64,
			Name: "Peter",
		},
		{
			Id:   1,
			Name: "Peter",
		},
		{
			1,
			"Jane",
		},
	}

	value, found := mapForIteration["First"]
	if !found {
		fmt.Println("Not found")
		return
	}
	fmt.Printf("%d ==> %v\n", value, found)

	uniqueUsers := make(map[int64]struct{}, len(students))

	for _, user := range students {
		if _, oK := uniqueUsers[user.Id]; !oK {
			uniqueUsers[user.Id] = struct{}{}
		}
	}

	fmt.Printf("Type %T value %#v\n", uniqueUsers, uniqueUsers)

	usersMap := make(map[int64]User, len(students))
	for _, v := range students {
		if _, ok := usersMap[v.Id]; !ok {
			usersMap[v.Id] = v
		}
	}

	fmt.Println(findInSlice(students, 1)) // find in slice return just 1 result, and do iterations 0-result
	fmt.Println(findInMap(usersMap, 1))
}

func findInSlice(users []User, id int64) *User {
	for _, user := range users {
		if user.Id == id {
			return &user
		}
	}
	return nil
}

func findInMap(usersMap map[int64]User, id int64) *User {
	if user, ok := usersMap[id]; ok {
		return &user
	}
	return nil
}
