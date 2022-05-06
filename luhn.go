package main

import (
	"fmt"
	"strconv"
)

func main() {
	var card string
	var sum int
	fmt.Println("Inserisci il numero della carta di credito:")
	fmt.Scanln(&card)

	for i := 0; i < len(card); i++ {
		intVal, _ := strconv.Atoi(string(card[i]))
		if intVal%2 != 0 {
			intVal = intVal * 2
			if len(strconv.Itoa(intVal)) >= 2 {
				str := strconv.Itoa(intVal)
				intVal1, _ := strconv.Atoi(string(str[0]))
				intVal2, _ := strconv.Atoi(string(str[1]))
				intVal = intVal1 + intVal2
			}
		}
		sum = sum + intVal
	}
	if sum%10 == 0 {
		fmt.Println("Valid card number")
	} else {
		fmt.Println("Invalid card number")
	}
}
