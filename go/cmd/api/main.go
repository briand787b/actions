package main

import (
	"fmt"

	"actions/internal/div"
	"actions/internal/sum"
)

func main() {
	fmt.Printf("4 + 1 = %d\n", sum.Sum(4, 1))
	fmt.Printf("4 / 1 = %d\n", div.CustomDivide(4, 1))
}
