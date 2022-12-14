package main

import "fmt"

func main() {

	var newSlice = make([]string, 2, 5) // len 2, cap 5
	newSlice[0] = "Otniel"
	newSlice[1] = "Kevin"

	fmt.Println(newSlice)      //[Otniel Kevin]
	fmt.Println(len(newSlice)) //2
	fmt.Println(cap(newSlice)) //5

	copySlice := make([]string, len(newSlice), cap(newSlice)) // len 2, cap 5
	copySlice2 := make([]string, 1, cap(newSlice))            // len dan cap nya harus sama, apabila tidak maka akan terpotong
	copy(copySlice, newSlice)
	copy(copySlice2, newSlice)

	fmt.Println(copySlice)  //[Otniel Kevin]
	fmt.Println(copySlice2) //[Otniel]

}
