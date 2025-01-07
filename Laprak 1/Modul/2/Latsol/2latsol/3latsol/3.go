package main

import "fmt"

func main() {
	var x, y, i int
	fmt.Scan(&x, &y)
	for i = 0; x >= y; i++ {
		x -= y
	}
	fmt.Print(i)
}
