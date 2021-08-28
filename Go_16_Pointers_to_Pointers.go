// Go - Pointers to Pointers

package main

import "fmt"

func main() {

	var a int = 10

	var p1 *int = &a
	var p2 **int = &p1
	var p3 ***int = &p2

	fmt.Println(a)
	fmt.Printf("%d\n", *p1)
	fmt.Printf("%d\n", **p2)
	fmt.Printf("%d\n", ***p3)

}
