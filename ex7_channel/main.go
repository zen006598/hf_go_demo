package main

import (
	"fmt"
	"sync"
	"time"
)

// `Channel` 是用來在不同 `goroutine` 之間安全地傳遞資料和進行同步的機制(共享資料)，
// 在並發中，直接共享記憶體容易導致競態條件和資料不一致的問題
// 「不要通過共享記憶體來通信，而是通過通信來共享記憶體」
// （"Don't communicate by sharing memory, share memory by communicating"）

func main() {
	// 不安全共享()
	// 安全共享()

	unbufferedChannel()
	// bufferedChannel()
	// readEmptyChannel()
}

func 不安全共享() {
	var counter int
	var wg sync.WaitGroup

	// 啟動 1000 個 goroutine 同時增加計數器
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			counter++
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("1000 /", counter)
}

func 安全共享() {
	var counter int
	ch := make(chan int, 1000)
	var wg sync.WaitGroup

	// 啟動 1000 個 goroutine 發送數字到 channel
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			ch <- 1 // 發送 1 到 channel
			wg.Done()
		}()
	}

	// 等待所有 goroutine 完成後關閉 channel
	wg.Wait()
	close(ch)

	// 從 channel 讀取並計數
	for n := range ch {
		counter += n
	}

	fmt.Printf("1000 / %d", counter)
}

// 一: 無緩衝 channel - 同步傳遞
// chan <-：阻塞，直到有另一個 goroutine 準備好從該 channel 讀取
// <- chan：阻塞，直到有另一個 goroutine 準備好向該 channel 寫入
func unbufferedChannel() {
	fmt.Println("=== 無緩衝 channel 測試 ===")
	ch := make(chan int) // 無緩衝 channel

	go func() {
		// 等待 3 秒後再接收，證明發送會被阻塞
		fmt.Println("in goroutine") // 1
		time.Sleep(3 * time.Second)
		fmt.Println("接收數據...") // 2
		value := <-ch
		fmt.Println("成功接收:", value) // 3
	}()

	// 發送操作 - 會阻塞直到有人接收
	fmt.Println("發送數據 42") // 1
	ch <- 42
	fmt.Println("發送成功!") // 3

	time.Sleep(time.Second)
}

// 二: Buffered channel - 非阻塞直到緩衝區滿
// 寫入：
// 當緩衝區未滿時：不阻塞
// 當緩衝區已滿時：阻塞，直到有空間

// 讀取：
// 當緩衝區有數據時：不阻塞
// 當緩衝區為空時：阻塞，直到有數據可用
func bufferedChannel() {
	fmt.Println("\n=== 有緩衝 channel 測試 ===")
	ch := make(chan int, 2) // 容量為 2 的緩衝 channel

	// 發送前兩個值不會阻塞
	fmt.Println("加入 1 channel")
	ch <- 1
	fmt.Println("加入 2 channel")
	ch <- 2
	fmt.Println("緩衝區已滿")

	// 嘗試發送第三個值 - 會阻塞，因為緩衝區已滿
	go func() {
		fmt.Println("發送 3")
		ch <- 3
		fmt.Println("3 發送成功!") // 只有在緩衝區有空間後才會執行
	}()

	// 等待 3 秒，讓上面的 goroutine 有時間阻塞
	time.Sleep(3 * time.Second)

	// 從 channel 讀取數據，釋放緩衝區空間
	fmt.Println("讀取數據 1:", <-ch)
	fmt.Println("讀取數據 2:", <-ch)
	fmt.Println("讀取數據 3:", <-ch)

	time.Sleep(100 * time.Millisecond)
}

// 三: 讀取空 buffered channel 會發生阻塞
func readEmptyChannel() {
	fmt.Println("\n=== 讀取空 channel 測試 ===")
	ch := make(chan int, 1)

	// 嘗試從空 channel 讀取 - 會阻塞
	go func() {
		fmt.Println("空 channel 讀取")
		value := <-ch
		fmt.Println("成功讀取:", value) // 只有在有數據可讀時才會執行
	}()

	// 等待一秒，讓上面的 goroutine 有時間阻塞
	time.Sleep(3 * time.Second)

	fmt.Println("發送數據到 channel")
	ch <- 100
	time.Sleep(100 * time.Millisecond)
}
