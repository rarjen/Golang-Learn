package main

import "fmt"

func main() {
	var name1 = "Kevin"
	var name2 = "Otniel"

	var result bool = name1 == name2
	var result2 = name1 > name2 //jarang digunakan membandingkan antar string

	fmt.Println(result)  //false
	fmt.Println(result2) //false, krn K lebih kecil daripada O, K lebih muncul lebih dulu

	var a = 2
	var b = 3

	var lebihdari = a > b
	fmt.Println(lebihdari) //false

	var kurangdari = a < b
	fmt.Println(kurangdari) //true

	var samadengan = a == b
	fmt.Println(samadengan) //false

	var tidakSama = a != b
	fmt.Println(tidakSama) //true
}
