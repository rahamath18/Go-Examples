// Go - Method
package main

import "fmt"

func main() {
	calc := MyCalculator{a: 10, b: 5}
	fmt.Println(calc.add())
	fmt.Println(calc.sub())
	fmt.Println(calc.multi())
	fmt.Println(calc.div())
}

type MyCalculator struct {
	a int
	b int
}

func (calc MyCalculator) add() int {
	return calc.a + calc.b
}

func (calc MyCalculator) sub() int {
	return calc.a - calc.b
}

func (calc MyCalculator) multi() int {
	return calc.a * calc.b
}

func (calc MyCalculator) div() int {
	return calc.a / calc.b
}
