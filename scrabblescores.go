package main

import "fmt"

func main() {
	var s string
	var p int
	fmt.Println("Inserisci una parola:")
	fmt.Scanln(&s)

	for i := 0; i < len(s); i++ {
		switch string(s[i]) {
		case "a", "e", "i", "o", "u", "l", "n", "r", "s", "t":
			p = p + 1
		case "d", "g":
			p = p + 2
		case "b", "c", "m", "p":
			p = p + 3
		case "f", "h", "v", "w", "y":
			p = p + 3
		case "k":
			p = p + 5
		case "j", "x":
			p = p + 8
		case "q", "z":
			p = p + 10
		}
	}
	fmt.Println(p)
}
