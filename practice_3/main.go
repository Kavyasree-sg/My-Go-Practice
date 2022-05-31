// convert string to other types
package main

import (
	"fmt"
	"math"
)

func main() {
	var aInt int = 20
	var aFloat float64 = 2.2
	// sum := aInt + aFloat // Not allowed and Go does not do implicit conversion
	sum := float64(aInt) + aFloat
	fmt.Println("The sum is ", sum)
	fmt.Println("The rounded sum is ", math.Round(sum))
}
