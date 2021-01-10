package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	c := 0

	fmt.Println("CPU:", runtime.NumCPU())
	fmt.Println("GOR:", runtime.NumGoroutine())
	fmt.Println("c  :", c)

	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)
	var mu sync.Mutex

	fmt.Println("---------------")

	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock()
			v := c
			runtime.Gosched()
			v++
			c = v
			mu.Unlock()
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
