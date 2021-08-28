// Go - Type Casting (conversion)

package main

import "fmt"

func main() {
	var total1 int = 3500
	var total2 int = 3500

	var grandTotal float64 = float64(total1) + float64(total2)

	fmt.Println(grandTotal)
}
