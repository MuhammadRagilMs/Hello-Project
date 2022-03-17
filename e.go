package main

import "fmt"

func main() {
	var kata string
	fmt.Scan(&kata)
	for kata != "selesai" {
		fmt.Print(kata)
		fmt.Scan(&kata)
	}
}