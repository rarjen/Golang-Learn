package main

import "fmt"

func main() {
	//Membuat dari awal tanpa ada data apapun
	//book := make(map[string]string)

	var book map[string]string = make(map[string]string)
	book["title"] = "Belajar Go-Lang"
	book["author"] = "Progammer Zaman Now"
	book["student"] = "Kevin"

	fmt.Println(book)

	//delete data
	delete(book, "student")
	fmt.Println(book)
}
