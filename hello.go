package main

import "fmt"

func main() {
	_, err := fmt.Println("Hello")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
}
