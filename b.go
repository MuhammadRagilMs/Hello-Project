package main

import "fmt"

func main() {
	var r int
	var luas_lingkaran float64

	fmt.Scanln(&r)
	luas_lingkaran = (22.0 / 7.0) * float64(r)
	fmt.Println("luas lingkaran dengan jari-jari =", r, "adalah", luas_lingkaran)
}