package main

import (
	"fmt"
	"math"
)

func main() {

	var floatNumber float64
	fmt.Scan(&floatNumber)

	truncated := int(math.Trunc(floatNumber))

	fmt.Println("Original number:", floatNumber)
	fmt.Println("Truncated number:", truncated)
}
