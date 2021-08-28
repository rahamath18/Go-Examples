// Go - Swap Pointers

package main

import "fmt"

func main() {

	var a int = 10
	var b int = 5

	fmt.Println(a, " | ", &a)
	fmt.Println(b, " | ", &b)

	var p1 *int = &a
	var p2 *int = &b

	var temp int

	temp = *p1

	*p1 = *p2
	*p2 = temp

	fmt.Println(a, " | ", &a)
	fmt.Println(b, " | ", &b)

	fmt.Println(p1)
	fmt.Println(p2)

}
