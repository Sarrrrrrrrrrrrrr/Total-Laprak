package main

import "fmt"

func main() {
	var x, i int
	fmt.Scan(&x)
	i = 0
	for x != 0 {
		l := x % 10
		i = i*10 + l
		x /= 10

	}
	fmt.Println(i, " ")
}
