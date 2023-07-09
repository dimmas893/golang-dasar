package main

import "fmt"

func main() {
	//const value yang tidak bisa di ubah
	const awalNama string = "ananda"
	const akhirNama = "dimmas"
	const value = 1000

	fmt.Println(awalNama)
	fmt.Println(akhirNama)
	fmt.Println(value)

	//cara lain
	const (
		awalNamaconst  string = "ananda"
		akhirNamaconst        = "dimmas"
		valueconst            = 1000
	)

	fmt.Println(awalNamaconst)
	fmt.Println(akhirNamaconst)
	fmt.Println(valueconst)
}
