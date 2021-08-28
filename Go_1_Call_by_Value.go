// Go - Call By Value
package main

import "fmt"

func main() {

	var one int = 10

	fmt.Println(one) // output = ?
	funcOne(one)
	fmt.Println(one) // output = ?

}
func funcOne(one int) {
	fmt.Println(one) // output = ?
	one = 11
	fmt.Println(one) // output = ?
}
