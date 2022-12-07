package main

import "fmt"

func main() {

	type NoKTP string // NoKTP adalah alias untuk string
	type Married bool // Married adalah alias untuk boolean
	type age int8     // age adalah alias untuk boolean
	// byte() //ini sebernanya adalah alias untuk uint8 (type byte = uint8)

	var ktpKevin NoKTP = "1111111111111" //menggunakan NoKTP dimana merupakan alias untuk string
	var marriedKevin Married = false
	var umur age = 21

	fmt.Println(ktpKevin) //1111111111111
	fmt.Println(NoKTP("2222222222222"))
	fmt.Println(umur)
	fmt.Println(marriedKevin) //false

}
