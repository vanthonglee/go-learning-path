package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	alex := person{firstName: "Alex", lastName: "Aderson"}

	fmt.Println(alex)

	var george person
	george.firstName = "George"
	george.lastName = "Lee"
	fmt.Println(george)
	fmt.Printf("%+v", george)

}
