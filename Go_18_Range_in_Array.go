// Go - Range in Array

package main

import "fmt"

func main() {
	valArr := []string{"Apple", "Banana", "Carret", "Donut", "EggPlant", "Orange", "Cucumber"}
	for i := range valArr {
		fmt.Print(valArr[i])
	}
}
