package main

import "fmt"

func main() {
	/*
		Hati - hati saat membuat array
		-> Saat membuat array, kita harus berhayi-hati, jk salah, maka yg kita buat bukanlah Array, tapi slice
	*/

	iniArray := [5]int{1, 2, 3, 4, 5} //array (harus ada panjangnya)
	iniSlice := []int{1, 2, 3, 4, 5}  //slice

	fmt.Println(iniArray) //[1 2 3 4 5]
	fmt.Println(iniSlice) //[1 2 3 4 5]
}
