package main

import (
	"fmt"
	"runtime"
	"sync"
)

func init() {
	fmt.Println("init")
}

var wg sync.WaitGroup

func main() {
	fmt.Println("Sequential------------------------------")
	fmt.Println("OS:\t\t\t\t", runtime.GOOS)
	fmt.Println("ARCH:\t\t\t", runtime.GOARCH)
	fmt.Println("CPUs:\t\t\t", runtime.NumCPU())
	fmt.Println("Go routines:\t", runtime.NumGoroutine())
	f1()
	f2()
	fmt.Println("OS:\t\t\t\t", runtime.GOOS)
	fmt.Println("ARCH:\t\t\t", runtime.GOARCH)
	fmt.Println("CPUs:\t\t\t", runtime.NumCPU())
	fmt.Println("Go routines:\t", runtime.NumGoroutine())
	fmt.Println("------------------------------")
	fmt.Println("OS:\t\t\t\t", runtime.GOOS)
	fmt.Println("ARCH:\t\t\t", runtime.GOARCH)
	fmt.Println("CPUs:\t\t\t", runtime.NumCPU())
	fmt.Println("Go routines:\t", runtime.NumGoroutine())
	wg.Add(1)
	go f3()
	f2()
	fmt.Println("OS:\t\t\t\t", runtime.GOOS)
	fmt.Println("ARCH:\t\t\t", runtime.GOARCH)
	fmt.Println("CPUs:\t\t\t", runtime.NumCPU())
	fmt.Println("Go routines:\t", runtime.NumGoroutine())
	wg.Wait()
}

func f1() {
	for i := 0; i < 10; i++ {
		fmt.Println("func1:", i)
	}
}

func f2() {
	for i := 0; i < 10; i++ {
		fmt.Println("func2:", i)
	}
}

//copy of f1 but with wg
func f3() {
	for i := 0; i < 10; i++ {
		fmt.Println("func1:", i)
	}
	wg.Done()
}
