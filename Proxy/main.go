package main

import (
	"fmt"
	"time"
)

func ProfileDecorator[T, R any](a func(T) R) func(T) R {

	return func(params T) R {
		defer func(t time.Time) {
			fmt.Printf("--- Time Elapsed 2: %v ---n", time.Since(t))
		}(time.Now())
		result := a(params)
		return result
	}
}

func main() {
	decoratedAdd := ProfileDecorator(duplicate)
	fmt.Println(decoratedAdd(2))
}

func duplicate(b int32) int32 {

	// time.Sleep(2 * time.Second)
	return b + b

}
