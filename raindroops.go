package main

import "fmt"

func main() {
	var i int
	var s string
	fmt.Println("Inserisci un numero:")
	fmt.Scan(&i)
	if i%3 == 0 {
		s += "Pling"
	}
	if i%5 == 0 {
		s += "Plang"
	}
	if i%7 == 0 {
		s += "Plong"
	}
	if len(s) == 0 {
		print(i)
	} else {
		fmt.Printf(s)
	}
}
