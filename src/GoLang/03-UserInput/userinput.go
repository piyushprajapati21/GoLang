package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to go lan"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter your ratings : ")
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for ratings : ", input)

	fmt.Println("Please confirm again your ratings : ")
	input1, _ := reader.ReadString('\n')
	fmt.Printf("Thanks for first %v and second %v ratings : ", input, input1)
}
