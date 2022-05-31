package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	d := rand.Intn(7) + 1
	fmt.Println("Day", d)
	var result string
	switch d {
	case 1:
		result = "It's Sunday"
	case 2:
		result = "It's Monday"
	case 3:
		result = "It's Tuesday"
	default:
		result = "It's some other day"
	}
	fmt.Println(result)
}
