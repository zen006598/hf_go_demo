package main

import "fmt"

func main() {
	arr()
	sli()
}

func arr() {
	// var x [3]int
	x := [3]int{5, 2, 3}
	// x := [5]int{5, 2, 3}
	// x := [5]int{2: 3, 40, 15}
	// [0, 0, 3, 40, 15]

	// 可以推斷長度
	// var x [...]int
	// var x [2][5]int = [2][5]int{{1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}}
	fmt.Println(x)
}

func sli(){
	x := []int{5, 2, 3}

	var ar = [10]byte {'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}
	var a, b []byte
	a = ar[2:5]
	b = ar[3:5]
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(x)
}