package main

import "fmt"

func main() {
	var fx, x float32
	fmt.Scan(&x)
	fx = (2 / (x + 5)) + (5)
	fmt.Println(fx)
}
