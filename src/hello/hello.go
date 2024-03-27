package main

import "fmt"

func main() {
	forLoop()
}

func forLoop() {
	for i := 1; i <= 5; i++ {
		fmt.Println("Number:", i)
	}
}
