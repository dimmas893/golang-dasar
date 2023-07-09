package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "ananda"
	names[1] = "dimmas"
	names[2] = "budiarto"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	//cara lain
	var values = [3]string{
		"ananda",
		"dimmas",
		"budiarto",
	}
	fmt.Println(values)

	fmt.Println(len(values)) //panjang array bukan jumlah isi data array
	fmt.Println(len(names))  //panjang array bukan jumlah isi data array
}
