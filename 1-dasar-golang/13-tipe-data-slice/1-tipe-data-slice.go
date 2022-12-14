package main

import "fmt"

func main() {
	var months = [12]string{
		"January",
		"February",
		"March",
		"April",
		"May",
		"June",
		"July",
		"August",
		"September",
		"October",
		"November",
		"December",
	}

	fmt.Println(months)
	fmt.Println(len(months)) //12

	// Tipe Data Slice
	var slice1 = months[4:7]
	fmt.Println(slice1)      //[May June July]
	fmt.Println(len(slice1)) //mengambil panjang slice (3)
	fmt.Println(cap(slice1)) //mengambil capasity (8) dihitung mulai dari index ke 4 - index terakhir (12)

	months[5] = "Diubah" //mengubah array months index ke 5
	fmt.Println(slice1)  //[May Diubah July] Di dalam slice pun ikut berubah datanya

	slice1[0] = "diubah2" //mengubah data slice index ke 0 dan otomatis mengubah array months index ke 4
	fmt.Println(months)   //January February March April diubah2 Diubah July August September October November December]

}
