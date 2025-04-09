package main

import (
	"fmt"
	"time"
)

func main() {
	// unbufferChannelIsLock()
	// runServer()
	run()
}

// unbufferChannel 沒有容量，當發送時會阻塞，直到有接收者接收數據
func unbufferChannelIsLock() {
	c := make(chan int)
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("發送 1")
		c <- 1
	}()

	fmt.Println("接收 1")
	<-c
}

// bufferChannel
const 最大併發數 = 3

var sem = make(chan int, 最大併發數)

func processSomething(id int) {
	sem <- 1
	fmt.Printf("工作 #%d 開始\n", id)

	time.Sleep(2 * time.Second)

	fmt.Printf("工作 #%d 完成\n", id)
	<-sem
}

func runServer() {
	fmt.Printf("開始執行，最多同時 %d 個工作\n", 最大併發數)

	for i := 1; i <= 7; i++ {
		go processSomething(i)
	}

	time.Sleep(6 * time.Second)
	fmt.Println("所有工作已完成")
}
