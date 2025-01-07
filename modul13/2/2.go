package main

import "fmt"

func main() {
	var number int
	var continueloop bool
	for continueloop = true; continueloop; {
		fmt.Scan(&number)
		continueloop = number <= 0
	}
	fmt.Printf("%d adalah bilangan bulat positif\n", number)
}
