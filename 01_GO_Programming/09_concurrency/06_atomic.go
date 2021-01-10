package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	var c int64

	fmt.Println("CPU:", runtime.NumCPU())
	fmt.Println("GOR:", runtime.NumGoroutine())
	fmt.Println("c  :", c)

	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	fmt.Println("---------------")

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&c, 1)
			fmt.Println("c=", atomic.LoadInt64(&c))
			runtime.Gosched()
			wg.Done()
		}()
		fmt.Println("GOR:", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("---------------")
	fmt.Println("CPU:", runtime.NumCPU())
	fmt.Println("GOR:", runtime.NumGoroutine())
	fmt.Println("c  :", c)
}
