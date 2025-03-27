package main

import "fmt"

func main() {
	// ex1()
	ex2_1()
	ex2_2()
}

func ex1(){
	num := 10             // 資料變數
	fmt.Println(num)      // output: 10
	fmt.Println(&num)     // &變數 取得記憶體位置, output: 0x1400000e1a0
	var place *int = &num // 變數指標 *type 存放記憶體位置
	fmt.Println(place)    // output: 0x1400000e1a0
	fmt.Println(*place)   // 反解記憶體位置 *變數名稱 output: 10
	*place = 20           // 反解記憶體位置並修改記憶體上的值
	fmt.Println(num)      // output: 20
}

func ex2_1(){
	num := 10
	num = ex2_1Add10(num)
	fmt.Println(num)
}

func ex2_1Add10(n int) int {
	n += 10
	return n
}

// side effect
func ex2_2(){
	num := 10
	ex2_2Add10(&num)
	fmt.Println(num)
}

func ex2_2Add10(n *int) {
	*n += 10
}