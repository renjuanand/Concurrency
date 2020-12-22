package main

import (
	"fmt"
)

func describe(t interface{}) {
	switch v := t.(type) {
	case int:
		fmt.Printf("integer: %d\n", v)
	case string:
		fmt.Printf("string: %s\n", v)
	default:
		fmt.Printf("unknown: %v\n", v)
	}
}

func main() {
	describe(10)
	describe("hello world")
	describe(123.45)
}
