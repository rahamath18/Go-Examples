// Go - Pointers in Array

package main

import "fmt"

func main() {
	var pArr [3]*string
	valArr := []string{"Apple", "Banana", "Carret"}
	var i int
	for i = 0; i < 3; i++ {
		pArr[i] = &valArr[i]
	}
	i = 0
	for i = 0; i < 3; i++ {
		fmt.Println(i, *pArr[i])
	}
}
