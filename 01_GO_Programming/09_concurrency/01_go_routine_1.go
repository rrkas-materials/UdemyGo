package main

import (
	"fmt"
	"runtime"
)

func init() {
	fmt.Println("init")
}

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

	fmt.Println("WITH CONCURRENCY------------------------------")
	fmt.Println("OS:\t\t\t\t", runtime.GOOS)
	fmt.Println("ARCH:\t\t\t", runtime.GOARCH)
	fmt.Println("CPUs:\t\t\t", runtime.NumCPU())
	fmt.Println("Go routines:\t", runtime.NumGoroutine())
	go f1()
	f2()
	fmt.Println("OS:\t\t\t\t", runtime.GOOS)
	fmt.Println("ARCH:\t\t\t", runtime.GOARCH)
	fmt.Println("CPUs:\t\t\t", runtime.NumCPU())
	fmt.Println("Go routines:\t", runtime.NumGoroutine())
	//NO func1 because even before a new routine is launched, main thread finishes its execution
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
