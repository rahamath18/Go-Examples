// Go - Map

package main

import "fmt"

func main() {

	fruitMap := map[string]string{"Apple": "5", "Banana": "12", "Carret": "2 KG", "Donut": "1 BOX", "EggPlant": "3 KG", "Orange": "2 Dozens", "Cucumber": "5 Pieces"}

	// key, value based
	for fruit, unit := range fruitMap {
		fmt.Println(fruit, unit)
	}
}
