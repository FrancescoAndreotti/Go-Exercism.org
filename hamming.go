package main

import "fmt"

func main() {
	s := "GAGCCTACTAACGGGAT"
	t := "CATCGTAATGACGGCCT"
	var c int

	if len(s) == len(t) {
		for i := 0; i < len(s); i++ {
			if s[i] != t[i] {
				c++
			}
		}
		fmt.Println("The Hamming Distance is:", c)
	}

}
