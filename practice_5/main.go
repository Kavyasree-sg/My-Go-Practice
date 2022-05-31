package main

import "fmt"

// an array is an object in go

func main() {
	var colors [3]string
	colors[0] = "red"
	colors[1] = "black"
	fmt.Println(colors)
	fmt.Println(colors[2])
	var numbers = [5]int{10, 20, 30, 40, 50}
	fmt.Println(numbers)

	// slices (resizing is easier)
	var colors1 = []string{"Dusty pink", "Green"}
	fmt.Println(colors1[0])
}
