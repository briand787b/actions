package main

import (
	"fmt"

	"actions/internal/div"
	"actions/internal/sum"
)

// this should obv be an API, but I'll keep it simple since this is a demo
func main() {
	fmt.Printf("4 + 1 = %d\n", sum.Sum(4, 1))
	fmt.Printf("4 / 1 = %d\n", div.CustomDivide(4, 1))
}
