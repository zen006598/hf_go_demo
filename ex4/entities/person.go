package entity

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

// 值型別，不需要原始值
func (p Person) GetAge() int {
	return p.Age
}

// 會需要原始值的話，就用指標
func (p *Person) SetAge(newAge int) {
	p.Age = newAge
}

func (p Person) Greet() {
	fmt.Println("Hello")
}

type Men struct {
	Person
	Beard bool
}
