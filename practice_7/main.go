package main

import "fmt"

func main() {
	colors := []string{"white", "black", "green", "blue"}
	// For loops
	for i := 0; i < len(colors); i++ {
		fmt.Println(colors[i])
	}
	fmt.Println("######################")
	for i := range colors {
		fmt.Println(colors[i])
	}
	fmt.Println("######################")
	for _, color := range colors {
		fmt.Println(color)
	}
	fmt.Println("######################")
	value := 1
	for value < 10 {
		fmt.Println("Value :", value)
		value++
	}
}
