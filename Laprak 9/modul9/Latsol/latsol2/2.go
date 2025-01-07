package main

import "fmt"

func main() {
	var x, y int
	fmt.Scan(&x)
	fmt.Scan(&y)

	fmt.Println(y%x == 0)
	fmt.Println(x%y == 0)
}
