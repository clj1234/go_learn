package main

import (
	"fmt"
	"time"
)

/*
1、题目 ：编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
考察点 ： go 关键字的使用、协程的并发执行。
2、题目 ：设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。
考察点 ：协程原理、并发任务调度。
*/
func main() {
	// problem1()
	// time.Sleep(2 * time.Second)
	problem2()
}

func problem1() {
	go func() {
		for i := 1; i <= 10; i++ {
			fmt.Println(i)
			time.Sleep(100 * time.Millisecond)
		}
	}()
	go func() {
		for i := 2; i <= 10; i++ {
			fmt.Println(i)
			time.Sleep(100 * time.Millisecond)
		}
	}()
}

/*
2、题目 ：设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。
考察点 ：协程原理、并发任务调度。
*/

type Task func()

func problem2() {
	ch1 := make(chan Task, 4)
	timeout := time.After(5 * time.Second)
	for i := 0; i < 3; i++ {
		go worker(ch1, timeout)
	}
	fmt.Println("准备执行提交任务")
	go func() {
		for i := 0; i < 10; i++ {
			taskId := i
			fmt.Println("提交任务：", i)
			ch1 <- func() {
				oldTime := time.Now().UnixMilli()
				fmt.Println("执行任务：", taskId)
				time.Sleep(time.Duration((i*100)+1000) * time.Millisecond)
				newTime := time.Now().UnixMilli()
				fmt.Printf("任务%d---usedTime:%d\n", i, newTime-oldTime)
			}
		}
	}()
	select {
	case <-timeout:
		fmt.Println("3秒后结束")
	}
	close(ch1)
	time.Sleep(1000 * time.Millisecond)
}

func worker(taskChan <-chan Task, timeout <-chan time.Time) {
	for {
		select {
		case task, ok := <-taskChan:
			if !ok {
				fmt.Println("channel已关闭")
				return
			}
			task()
		case <-timeout:
			fmt.Println("超时停止")
			return
		}
	}
}
