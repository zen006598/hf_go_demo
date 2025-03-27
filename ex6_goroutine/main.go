package main

import (
	"fmt"
	"sync"
	"time"
)

// goroutine

func main() {
	// ex1()
	// ex2()
	// ex3()
	// ex4()
}

func ex1() {
	// 計算 1-10
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(i)
			time.Sleep(100 * time.Millisecond)
		}
	}()
	// 計算 11-20
	go func() {
		for i := 11; i <= 20; i++ {
			fmt.Println(i)
			time.Sleep(100 * time.Millisecond)
		}
	}()

	fmt.Println("start")

	// 等待 count() 執行完畢
	time.Sleep(1 * time.Second)
	fmt.Println("end")
}

// 父程式結束，其中的子攜程也會關閉
func ex2() {
	fmt.Println("start")
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(i)
			time.Sleep(100 * time.Millisecond)
		}
	}()
	fmt.Println("end")
}

// 等待 Goroutine 執行完才結束程式的方法
// 1. 使用 time.Sleep() 來等待
// 2. 使用 WaitGroup 來等待
// 3. 使用 channel 來等待

// waitGroup
// Add() / Done() / Wait()
func ex3() {
	var wg sync.WaitGroup

	// 啟動 5 個 goroutine
	for i := 1; i <= 5; i++ {
		wg.Add(1) // 通知 WaitGroup 多一個任務
		go worker(i, &wg)
	}

	// 等待所有 goroutine 完成
	wg.Wait()
	fmt.Println("所有工作者都完成了。")
}

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // 表示一定會執行 通知 WaitGroup 這段已經完成
	fmt.Printf("Worker %d 正在工作...\n", id)
}

func ex4() {
	c := make(chan int)
	go doSomething(c)
	<-c
}

func doSomething(c chan int) {
	fmt.Println("test")
	c <- 1
}
