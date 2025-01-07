package main

import "fmt"

func main() {
	var n, count int
	fmt.Scan(&n)
	for count = 0; n > 0; count++ {
		n /= 10
	}
	fmt.Println(count)
}
