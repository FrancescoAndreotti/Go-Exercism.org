package main

import "fmt"

func main() {
	nums := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var a int
	var b int
	var d int
	for _, s := range nums {
		a = a + s*s
		b = b + s
	}
	d = (b * b) - a
	fmt.Println(d)

}
