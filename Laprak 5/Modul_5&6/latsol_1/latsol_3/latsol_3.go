package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a)
	fmt.Scan(&b)
	result := 1
	for i := 0; i < b; i++ {
		result *= a
	}

	fmt.Print(result)
}
