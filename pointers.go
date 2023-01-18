package main

import "fmt"

type zooInfo struct {
	town    string
	email   string
	pinCode int
}

type animal struct {
	name  string
	sound string
	zooInfo
}

func main() {

	cat := animal{
		name:  "Cat",
		sound: "Meeeaaawwwwww",
		zooInfo: zooInfo{
			town:    "California",
			email:   "abc@xyz.com",
			pinCode: 12345,
		},
	}

	cat.print("Original Structure")
	catPointer := &cat
	catPointer.updateAttributes("Dog", "Woof Woof Woof")

	cat.print("Updated Struct")
}

func (p *animal) updateAttributes(newName string, sound string) {
	(*p).name = newName
	(*p).sound = sound
}

func (p animal) print(text string) {
	result := fmt.Sprintf("%s ==> %+v \n \n", text, p)
	fmt.Printf("%+v", result)
}
