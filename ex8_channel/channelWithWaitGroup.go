package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// 建立全局亂數產生器
var rng = rand.New(rand.NewSource(time.Now().UnixNano()))

// 工作函數：計算一個數字的平方，但需要一些時間
func calculateSquare(number int, results chan<- int, wg *sync.WaitGroup) {
	// 完成時通知 WaitGroup
	defer wg.Done()

	// 模擬工作需要隨機時間
	delay := time.Duration(1+rng.Intn(3)) * time.Second
	fmt.Printf("工作 #%d 開始計算，需要 %v\n", number, delay)
	time.Sleep(delay)

	// 計算平方並送入結果通道
	result := number * number
	fmt.Printf("工作 #%d 完成計算: %d 的平方是 %d\n", number, number, result)

	// 將結果送入通道
	results <- result
}

func run() {
	var wg sync.WaitGroup
	results := make(chan int, 5)
	fmt.Println("start jobs:")

	// 啟動 5 個工作
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go calculateSquare(i, results, &wg)
	}

	// 啟動一個 goroutine 來等待所有工作完成並關閉結果通道
	go func() {
		wg.Wait()
		fmt.Println("所有工作已完成，關閉 channel")
		close(results)
	}()

	// 主 goroutine 處理所有的結果
	fmt.Println("開始收集結果:")

	sum := 0
	// 從通道讀取直到它關閉
	for result := range results {
		fmt.Printf("channel 收到結果: %d\n", result)
		sum += result
	}

	fmt.Printf("所有結果的總和: %d\n", sum)
}
