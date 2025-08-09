package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
Channel
1、题目 ：编写一个程序，使用通道实现两个协程之间的通信。一个协程生成从1到10的整数，并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来。
考察点 ：通道的基本使用、协程间通信。
2、题目 ：实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。
考察点 ：通道的缓冲机制。
*/
func main() {
	taskChan := make(chan int, 10)

	timeout := 3 * time.Second
	timer := time.NewTimer(timeout)
	var wg sync.WaitGroup
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go worker(taskChan, timer, &wg)
	}
	// 题目1
	go producer(taskChan)
	// 题目2
	go func() {
		for i := 0; i < 100; i++ {
			taskChan <- rand.Int()
		}
	}()
	wg.Wait()
	fmt.Println("主线程结束")
}

type Task func()

func producer(taskChan chan<- int) {
	for i := 1; i <= 10; i++ {
		taskChan <- i
	}

}

func worker(taskChan chan int, timer *time.Timer, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case num, ok := <-taskChan:
			if !ok {
				fmt.Println("channel已关闭")
				return
			}
			fmt.Println("receiveNum :", num)
		case <-timer.C:
			fmt.Println("所有任务处理完成")
			close(taskChan)
			return
		}
	}
}
