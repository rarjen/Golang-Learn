package main

import "fmt"

func main() {
	var nilai32 int32 = 100000
	var nilai64 int64 = int64(nilai32) // konversi dari 32 ke 64
	var nilai8 int8 = int8(nilai32)

	fmt.Println(nilai32) //10000
	fmt.Println(nilai64) //100000
	fmt.Println(nilai8)  //-96 krn tidak bisa handle angka diatas 128
	// terjadi integer overflow dimana jika sudah mencapai batas akan kembali dari paling bawah

}
