package main

import "fmt"

func main() {
	var values = [3]int{
		11,
		22,
		33,
	}

	//reprentasi string dari array
	fmt.Println(values) //[11 22 33]

	fmt.Println(values[0]) //11
	fmt.Println(values[1]) //22
	fmt.Println(values[2]) //33

	var len = len(values)
	fmt.Println(len) //3

	// var lagi [10]string

	// fmt.Println(len(lagi)) // 10

}
