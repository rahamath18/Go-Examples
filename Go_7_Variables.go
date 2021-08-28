// Go - Multiple types of variables in single declaration
package main

import "fmt"

func main() {

	// multiple types of variables in single declaration in Go
	var a, b, c = 3, 4, "foo"

	fmt.Println(a) // output = ?
	fmt.Println(b) // output = ?
	fmt.Println(c) // output = ?

	// var name="Anil", city="BLR" - compilatoin error
	var name, city = "Anil", "BLR"
	fmt.Println(name) // output = ?
	fmt.Println(city) // output = ?

}
