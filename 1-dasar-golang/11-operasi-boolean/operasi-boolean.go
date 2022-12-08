package main

import "fmt"

func main() {
	var nilaiAkhir = 90
	var absen = 75

	var lulusNilaiAkhir = nilaiAkhir >= 65
	var lulusAbsensi = absen >= 75

	var lulus = lulusNilaiAkhir && lulusAbsensi

	fmt.Println(lulus) //true
}
