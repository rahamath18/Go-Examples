// Go - function closure
package main

import "fmt"

func nextValue() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func main() {
	nextNumber := nextValue()

	fmt.Println(nextNumber()) // output = ?
	fmt.Println(nextNumber()) // output = ?
	fmt.Println(nextNumber()) // output = ?

	nextNumber1 := nextValue()
	fmt.Println(nextNumber1()) // output = ?
	fmt.Println(nextNumber1()) // output = ?
}
