package main

import "fmt"

func main() {
	fmt.Println(len("Otniel"))           // Panjang String (6 = adalah  representasi byte dari O)
	fmt.Println("Otniel Kevin"[0])       // mengambil karakter pertama (79 = adalah representasi byte dari O)
	fmt.Println("Otniel Kevin Abiel"[1]) // mengambil karakter kedua (116 = adalah  representasi byte dari O)
}
