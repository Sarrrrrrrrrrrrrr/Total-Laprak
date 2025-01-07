package main

import "fmt"

func main() {
	var token string
	for token != "12345abcde" {
		fmt.Scan(&token)
	}
	fmt.Println("selamat anda berhasil login")
}
