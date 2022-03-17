package main
import "fmt"

func main() {
	var S string
	var A, B, hasil_penjumlahan int

	fmt.Scan(&S, &A, &B)
	hasil_penjumlahan = A + B
	fmt.Print("kata = ", S)
	fmt.Print("Jumlah = ", hasil_penjumlahan)
}