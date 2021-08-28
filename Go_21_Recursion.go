// Go - Recursion

package main

import "fmt"

func factorial(i int) int {
	fmt.Println("->", i)
	if i <= 1 {
		return 1
	}
	return i * factorial(i-1)
}
func main() {
	var i int = 5
	fmt.Printf("Factorial of %d is %d", i, factorial(i))
}
