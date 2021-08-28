// := short variable declaration
package main

import "fmt"

func main() {

	name, location, awards := func_With_Multi_Returns("Player ")
	fmt.Println(name, location, awards) // output = ?

}

func func_With_Multi_Returns(b string) (string, string, int) {
	return b + "Jonn", "From UK comes in place of", 100
}
