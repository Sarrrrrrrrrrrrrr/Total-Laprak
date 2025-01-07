package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	result := 0
	for a := 1; a <= n; a++ {
		result += a

	}
	fmt.Print(result)
}
