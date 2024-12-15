package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func NewPerson(name string, age int) Person {
	return Person{
		Name: name,
		Age:  age,
	}

}
func (p *Person) Change_name(name string) {
	p.Name = name
}
func main() {

	my_person := NewPerson("rajesh tope", 12)
	fmt.Printf("this is my person %+v \n", my_person)
	my_person.Change_name("gote")
	fmt.Printf("this is my person %+v \n", my_person)

}
