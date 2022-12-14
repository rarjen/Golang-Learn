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

	var slice2 = months[10:]
	// pointer : 10
	// cap : 10
	// len : 2
	fmt.Println(slice2)      //November December
	fmt.Println(len(slice2)) //2
	fmt.Println(cap(slice2)) //2

	var slice3 = append(slice2, "Leinto", "1", "2", "3", "4", "5", "6", "8", "9", "10", "11")
	fmt.Println(slice3) //November Desember Leinto (Saat Kita append dan slice2 sudah penuh capasities nya, maka buat array baru. slice3 disini merupakan array baru)
	slice3[1] = "NotDecember"
	fmt.Println(slice3) // November NotDecember Leinto

	fmt.Println(slice2) // Tidak Berubah karena berada diluar cakupkan array
	fmt.Println(months) // Tidak Berubah karena berada diluar cakupkan array
}
