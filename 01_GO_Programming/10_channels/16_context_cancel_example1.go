package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()

	fmt.Println("context:\t\t", ctx)
	fmt.Println("context err:\t", ctx.Err())
	fmt.Printf("context type:\t %T\n", ctx)
	fmt.Println("----------")

	ctx2, _ := context.WithCancel(ctx)

	fmt.Println("context:\t\t", ctx2)
	fmt.Println("context err:\t", ctx2.Err())
	fmt.Printf("context type:\t %T\n", ctx2)

	fmt.Println("----------")

	fmt.Println("context:\t\t", ctx, "\t\t", ctx2)
	ctx2, ctx = ctx, ctx2 //can be assigned to each other
	fmt.Println("context:\t\t", ctx, "\t\t", ctx2)
}
