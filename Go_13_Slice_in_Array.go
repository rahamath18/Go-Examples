// Go - Slice in Array
package main

import "fmt"

func main() {
	valArr := []string{"Apple", "Banana", "Carret", "Donut", "EggPlant", "Orange", "Cucumber"}
	fmt.Println(len(valArr), cap(valArr), valArr)
	fmt.Println(valArr)
	fmt.Println(valArr[1])
	fmt.Println(valArr[0:3])
	fmt.Println(valArr[:3])
	fmt.Println(valArr[3:])

	sliceArr := valArr[0:5]
	for i := 0; i < len(sliceArr); i++ {
		fmt.Print(sliceArr[i], " ")
	}

}
