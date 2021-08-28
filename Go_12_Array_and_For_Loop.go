// Go - Array & For Loop
package main

import "fmt"

func main() {

	var fruitArr [3]string
	fruitArr[0] = "Apple"
	fruitArr[1] = "Banana"
	fruitArr[2] = "Carret"
	var i int
	fmt.Println("Array Size: ", len(fruitArr))
	for i = 0; i < 3; i++ {
		fmt.Println((i + 1), fruitArr[i], " | ", len(fruitArr[i]))
	}

}
