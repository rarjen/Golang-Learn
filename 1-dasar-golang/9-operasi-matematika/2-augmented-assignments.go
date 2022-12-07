package main

import "fmt"

func main() {
	var a = 20
	a += 10
	fmt.Println(a) //30

	var b = 20
	b -= 10
	fmt.Println(b) //10

	var z = 2
	var c = 3
	c *= z // c = c * z
	fmt.Println(c)

	var d = 10
	d /= 10        // d = d / 10
	fmt.Println(d) //1

	var e = 20
	e %= 7       // e = e % 10
	fmt.Print(e) //6

}
