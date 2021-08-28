// Go - Pointers
/*
What is the data type of a pointer ?
	The actual data type of the value of all pointers, whether integer, float, or otherwise,
	is the same, a long hexadecimal number that represents a memory address.
*/
// var age *int - Pointer
// &age - address
// age - actual value

package main

import "fmt"

func main() {
	var ip *int
	var fp *float32

	fmt.Println(ip) // output = ?
	fmt.Println(fp) // output = ?

	fmt.Printf("%x\n", fp) // output = ?

	var age int = 10
	var sal float32 = 10.12

	ip = &age
	fp = &sal

	fmt.Println(ip) // output = ?
	fmt.Println(fp) // output = ?

	fmt.Println(age) // output = ?
	fmt.Println(sal) // output = ?

}
