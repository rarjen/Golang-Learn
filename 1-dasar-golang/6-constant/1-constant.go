package main

import "fmt"

func main() {

	const firstName string = "Otniel"
	const lastName = "Kevin"
	const age = 21 // walaupun const tidak digunakan maka tidak akan error

	fmt.Println(firstName)
	fmt.Println(lastName)

	// firstName = "leinto" //Error tidak bisa diubah
	// lastName = "777" //Error tidak bisa diubah

}
