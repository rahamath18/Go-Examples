// Go - Call By Reference
package main

import "fmt"

func main() {

	var two int = 20
	fmt.Println(two) // output = ?
	funcTwo(&two)
	fmt.Println(two) // output = ?

}

func funcTwo(two *int) {
	fmt.Println(two) // output = ?
	*two = 21
	fmt.Println(two) // output = ?
}
