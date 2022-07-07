package main

import (
	"fmt"
	"time"
)

// the decorator function
func ProfileDecorator[T, R any](a func(T) R) interface{} {
	var params T
	defer func(t time.Time) {
		fmt.Printf("--- Time Elapsed: %v ---n", time.Since(t))
	}(time.Now())
	return a(params)
}

func main() {
	decoratedAdd := ProfileDecorator(duplicate)

}

func duplicate(b int32) int32 {
	return b + b
}

// type IntOrString interface {int32 | string}
// func duplicate( b IntOrString) IntOrString {
// 	return  b + b
// }
