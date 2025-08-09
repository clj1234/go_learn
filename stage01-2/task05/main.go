package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/*
锁机制
1、题目 ：编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
考察点 ： sync.Mutex 的使用、并发数据安全。
2、题目 ：使用原子操作（ sync/atomic 包）实现一个无锁的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
考察点 ：原子操作、并发数据安全。
*/

func main() {
	var wg sync.WaitGroup
	// 题目1
	var lock sync.Mutex
	count := 0
	problem1(&lock, &wg, &count)
	wg.Wait()
	fmt.Println("count:", count)
	// 题目2
	// var count atomic.Int32
	// problem2(&wg, &count)
	// wg.Wait()
	// fmt.Printf("count---atomic: %v\n", count.Load())
}

func problem1(lock *sync.Mutex, wg *sync.WaitGroup, count *int) {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 1000; i++ {
				lock.Lock()
				*count++
				lock.Unlock()
			}
		}()
	}
}

/*
2、题目 ：使用原子操作（ sync/atomic 包）实现一个无锁的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
考察点 ：原子操作、并发数据安全。
*/
func problem2(wg *sync.WaitGroup, count *atomic.Int32) {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 1000; i++ {
				count.Add(1)
			}
		}()
	}
}
