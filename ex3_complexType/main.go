package main

// 複合型態
import "fmt"

func main() {
	// arr()
	// sli()
	// sli2()
	// mp()
}

func arr() {
	// var 變數名稱[array 長度]資料型態
	// 初始化為 0
	var x [3]int

	// x := [3]int{5, 2, 3}

	// 陣列長度不足時，會補 0
	// x := [5]int{5, 2, 3}

	// [0, 0, 3, 40, 15]
	// 指定 index 賦值
	// x := [5]int{2: 3, 40, 15}

	// [...] 使用推斷長度
	// var x [...]int{1, 2, 3, 4, 5}

	// 二維陣列
	// var x [2][5]int = [2][5]int{{1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}}

	// 陣列比較
	// a := [3]int{1, 2, 3}
	// b := [3]int{1, 2, 3}
	// x := a == b
	fmt.Println(x)
}

func sli() {
	// [] 沒有長度 即為 slice
	// 初始為 nil
	x := []int{}
	// x := []int{5, 2, 3}
	fmt.Println(x)

	// 切割 slice
	var ar = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}
	a := ar[2:5]
	// ar[2] = "1"
	// b := ar[3:5]
	fmt.Println(a)
	// fmt.Println(b)

	// copy(目的, 來源)
	// b := make([]string, 2)
	// copy(b, ar)
	// fmt.Println(b)

}

// 儲存在 連續記憶體，每個 slice 都有容量上限，當達到上限，就會把 slice 複製到更大的新位置
// GC 回收舊 slice，所以如果知道大概有多少量就可以先指定容量 (make)
// cap 檢查容量
func sli2() {
	// 長度 5, 容量 5
	// x := make([]int, 5)
	// 長度 5, 容量 10
	x := make([]int, 5, 10)
	// [0 0 0 0 0]

	fmt.Println(x)
	fmt.Println(cap(x))
}

func mp() {
	// number := map[key]value{}
	// var number map[string]int
	// number := map[string]int{
	// 	"一": 1,
	// 	"二": 2,
	// }

	number := make(map[string]int)
	number["三"] = 3

	// 如果對於鍵值的數量有概念的話，也可以給定初始大小，將有助於效能提升（The little Go book）
	// m := make(map[string]int, 100)
	fmt.Println(number)
	fmt.Println(number["三"])
	// index 不存在時，會回傳 0
	// ok
	val, ok := number["三"]
	fmt.Println(val, ok)
	val2, ok := number["四"]
	fmt.Println(val2, ok)
}
