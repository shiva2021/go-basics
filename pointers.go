package main

import "fmt"

type contactInfo struct {
	email   string
	pinCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {

	alex := person{
		firstName: "Alex",
		lastName:  "Anderson",
		contactInfo: contactInfo{
			email:   "abc@xyz.com",
			pinCode: 12345,
		},
	}

	alex.print("Original Structure")
	alexPointer := &alex
	alexPointer.updateFirstName("Tom")

	alex.print("Updated Struct")
}

func (p *person) updateFirstName(newName string) {
	(*p).firstName = newName
}

func (p person) print(text string) {
	result := fmt.Sprintf("%s ==> %+v \n \n", text, p)
	fmt.Printf("%+v", result)
}
