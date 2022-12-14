package main

import "fmt"

func main() {
	var names [3]string //array yg menampung string dengan jumlah data max 3

	names[0] = "Otniel" //Input to index 0
	names[1] = "Kevin"  //Input to index 1
	names[2] = "Abiel"  //Input to index 2
	// names[3] = "Abiel"  //Error: Out of bounds (array overflow)

	//Akses data by index
	fmt.Println(names[0]) //Otniel
	fmt.Println(names[1]) //Kevin
	fmt.Println(names[2]) //Abiel

	var numbers [3]int

	numbers[0] = 1
	numbers[1] = 2
	numbers[2] = 3

	fmt.Println(numbers[0]) //1
	fmt.Println(numbers[1]) //2
	fmt.Println(numbers[2]) //3

}
