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
	fmt.Println("---------------")
	for i := 0; i < gs; i++ {
		go func() {
			v := c
			runtime.Gosched()
			v++
			c = v
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

//go run -race 04_race_condition.go
