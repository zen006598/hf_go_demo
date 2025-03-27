package main

import (
	"fmt"

	entity "github.com/zen006598/goDemo/ex4_struct/entities"
	"github.com/zen006598/goDemo/ex4_struct/enums"
)

// struct
type Point struct {
	X int
	Y int
}

func main() {
	// ex1()
	// ex2_值型別()
	// ex3()
	// ex4()
	ex5()
}

func ex1() {
	p := Point{}
	p.X = 3
	p.Y = 4
	println(p.X, p.Y)
}

func ex2_值型別() {
	p1 := Point{1, 2}
	p2 := p1

	p2.X = 10
	println(p1.X)
	println(p2.X)
}

// OOP
func ex3() {
	p := entity.Person{Name: "jeff", Age: 55}
	fmt.Println(p.GetAge())

	p.SetAge(66)
	fmt.Println(p.GetAge())

	p2 := entity.NewPerson("jeff2", 77)
	fmt.Println(p2.GetAge())
}

// 匿名欄位/嵌入式欄位 （類似繼承）
func ex4() {
	m1 := entity.Men{
		Person: entity.Person{
			Name: "jeff",
			Age:  55,
		},
		Beard: true,
	}
	fmt.Println(m1.GetAge())
	m1.Greet()
}

func ex5() {
	b := entity.Box{1, 2, 3, enums.RED}
	fmt.Println(b.Volume())
	b.SetColor(enums.BLUE)
	fmt.Println(b.Color)
	fmt.Println(b.Color.ToString())
}
