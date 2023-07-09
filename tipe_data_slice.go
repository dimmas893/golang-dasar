package main

import "fmt"

func main() {
	var months = [...]string{
		"januari",
		"februari",
		"maret",
		"april",
		"mei",
		"juni",
		"juli",
		"agustus",
		"september",
		"oktober",
		"november",
		"desember",
	}

	var slice1 = months[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	var slice2 = months[10:]
	fmt.Println(slice2)
	var slice3 = append(slice2, "ananda") //menambahkan data array jika kapasitas penuh, kalau array masih cukup maka di tambah
	fmt.Println(slice3)
	slice1[1] = "bukan desember"
	fmt.Println(slice3)
	fmt.Println(slice2)
	fmt.Println(months)

	//contoh penulisan
	iniArray := [...]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}
	fmt.Println(iniArray)
	fmt.Println(iniSlice)
	//batas contoh
 
	newslice := make([]string, 2, 5)
	newslice[0] = "ananda"
	newslice[1] = "dimmas"

	fmt.Println(newslice)
	fmt.Println(len(newslice))
	fmt.Println(cap(newslice))

	copySlice := make([]string, len(newslice), cap(newslice))
	copy(copySlice, newslice)
	fmt.Println(copySlice)
}
