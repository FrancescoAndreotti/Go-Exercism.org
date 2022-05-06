package main

import "fmt"

func main() {
	fmt.Println("Enter your name:")
	var name string
	fmt.Scanln(&name)

	if len(name) != 0 {
		fmt.Printf("One for %s, one for me.", name)
	} else {
		fmt.Println("One for you, one for me.")
	}
}
