package main

import (
	"fmt"
	"reflect"

	"github.com/zen006598/goDemo/ex5_interface/entities"
)

// interface
func main() {
	// ex1()
	// ex2()
	// ex3()
	// ex4()
}

func ex1() {
	rect := entities.Rect{Width: 10, Height: 20}
	circle := entities.Circle{Radius: 10}

	shapes := []entities.Shape{rect, circle}

	for _, shape := range shapes {
		fmt.Println(shape.Area())
	}
}

// 通用容器
func ex2() {
	var data []interface{}
	data = append(data, 42)
	data = append(data, "Hello, Go!")
	data = append(data, 3.14)
	data = append(data, true)

	for _, v := range data {
		switch value := v.(type) {
		case int:
			fmt.Printf("整數: %d\n", value)
		case string:
			fmt.Printf("字串: %s\n", value)
		case float64:
			fmt.Printf("浮點數: %f\n", value)
		case bool:
			fmt.Printf("布林值: %v\n", value)
		default:
			fmt.Printf("未知類型: %T\n", value)
		}
	}
}

// T 泛型
func ex3() {
	PrintValue(100)
	PrintValue("Gopher")
	PrintValue(3.14)
	PrintValue(false)
}

func PrintValue(v interface{}) {
	fmt.Printf("值: %v, 型別: %T\n", v, v)
}

// reflect
func ex4() {
	var x interface{} = "Hello, Reflection!"
	fmt.Printf("動態型別: %s, 值: %v\n",
		reflect.TypeOf(x),
		reflect.ValueOf(x))
}
