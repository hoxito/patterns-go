package main

import (
	"fmt"
	"time"
)

type Function func[T,R any]( v T) R

// the decorator function
func ProfileDecorator(fn Function) Function {
	return func(params ...interface{}) interface{} {
		defer func(t time.Time) {
			fmt.Printf("--- Time Elapsed: %v ---n", time.Since(t))
		}(time.Now())
		return fn(params)
	}
}

func main() {
	decoratedAdd := ProfileDecorator(add2)
}
func add2(a, b int32) int32 {
	return a + b
}
