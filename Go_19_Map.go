// Go - Map

package main

import "fmt"

func main() {
	var fruitMap map[string]string
	fruitMap = make(map[string]string)

	fruitMap["Apple"] = "5"
	fruitMap["Banana"] = "12"
	fruitMap["Carret"] = "2 KG"
	fruitMap["Donut"] = "1 BOX"
	fruitMap["EggPlant"] = "3 KG"
	fruitMap["Orange"] = "2 Dozens"
	fruitMap["Cucumber"] = "5 Pieces"

	for fruit := range fruitMap {
		fmt.Println(fruit, fruitMap[fruit])
	}
}
