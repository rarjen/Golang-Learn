package main

import "fmt"

func main() {

	// Make New Slice

	var newSlice = make([]string, 2, 5) // len 2, cap 5
	newSlice[0] = "Otniel"
	newSlice[1] = "Kevin"
	// newSlice[2] = "Abiel" //Error: out of range

	fmt.Println(newSlice)      //[Otniel Kevin]
	fmt.Println(len(newSlice)) //2
	fmt.Println(cap(newSlice)) //5

	//Catatan:
	/*
		Membuat Slice dari awal lebih bagus menggunakan :
		Ini jadi lebih aman krn array berada di dlm dan slice nya tidak keliatan krn tdk disimpan dalam variable seperti months
	*/
}
