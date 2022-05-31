package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your name : ")
	input, _ := reader.ReadString('\n')
	fmt.Print("You entered ", input)
}
