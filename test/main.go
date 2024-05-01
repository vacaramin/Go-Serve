package main

import "fmt"

type Person struct {
	name string
	age  int32
}

func newPerson(name string, age int32) Person {
	return Person{
		name: name,
		age:  age,
	}
}
func changeName(person *Person, name string) {
	person.name = name
}

func main() {
	person := newPerson("Waqar", 23)
	changeName(&person, "blue")
	fmt.Println(person)
}
