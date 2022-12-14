package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "Gilbert",
		"address": "Semarang",
		"gender":  "Male",
	}

	// Add Data to Map:
	person["status"] = "Single"

	fmt.Println(person)

	//Access by key
	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person["gender"])
	fmt.Println(person["status"])
	// fmt.Println(person["religion"])
	fmt.Println(len(person)) //4
}
