// struct
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello Struct!")

	fmt.Println(Person{"ames", 20})

	fmt.Println(Person{name: "james2", age: 34})

	var p1 Person
	var p2 Person
	p1.name = "wangdabao"
	p2.name = "chenchao"
	p1.age = 27
	p2.age = 27

	fmt.Printf("p1 name = %s ,age= %d\n", p1.name, p1.age)

	call_struct(p2)

	var pointer_person *Person

	pointer_person = &p1

	call_struct_pointer(pointer_person)
}

type Person struct {
	name string
	age  int
}

func call_struct(person Person) {

	fmt.Printf("call_struct person name = %s ,age= %d\n", person.name, person.age)

}

func call_struct_pointer(person *Person) {

	fmt.Printf("call_struct_pointer person name = %s ,age= %d\n", person.name, person.age)
}
