package main

import "fmt"

func main() {

	var name = "Otniel Kevin Abiel"
	var e byte = name[0] //Mengambil karakter pertama, byte = uint8
	// var e = name[0]    //Mengambil karakter pertama, byte = uint8
	var eString = string(e) //Konversi dari byte ke string

	fmt.Println(name)    // Otniel Kevin Abiel -> String
	fmt.Println(e)       //79 -> byte
	fmt.Println(eString) //O -> string

	var address = "Semarang"
	fmt.Println(string(address[0])) //S -> String

}
