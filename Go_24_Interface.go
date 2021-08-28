// Go - Type Casting (conversion)

package main

import "fmt"

type BasicMaths interface {
	add() int
	sub() int
	mul() int
	div() int
}

func getAdd(basic BasicMaths) int {
	return basic.add()
}

func getSub(basic BasicMaths) int {
	return basic.sub()
}

func getMul(basic BasicMaths) int {
	return basic.mul()
}

func getDiv(basic BasicMaths) int {
	return basic.div()
}

// BasicCalculator starts

type BasicCalculator struct {
	a int
	b int
}

func (calc BasicCalculator) add() int {
	return calc.a + calc.b
}

func (calc BasicCalculator) sub() int {
	return calc.a - calc.b
}

func (calc BasicCalculator) mul() int {
	return calc.a * calc.b
}

func (calc BasicCalculator) div() int {
	return calc.a / calc.b
}

// ScientificCalculator starts

type ScientificCalculator struct {
	a int
	b int
}

func (calc ScientificCalculator) add() int {
	return calc.a + calc.b
}

func (calc ScientificCalculator) sub() int {
	return calc.a - calc.b
}

func (calc ScientificCalculator) mul() int {
	return calc.a * calc.b
}

func (calc ScientificCalculator) div() int {
	return calc.a / calc.b
}

func main() {
	fmt.Println("Basic Calculator")
	basicCalc := BasicCalculator{a: 10, b: 5}
	fmt.Println(getAdd(basicCalc))
	fmt.Println(getSub(basicCalc))
	fmt.Println(getMul(basicCalc))
	fmt.Println(getDiv(basicCalc))

	fmt.Println("Scientific Calculator")
	scientificCalc := ScientificCalculator{a: 25, b: 4}
	fmt.Println(getAdd(scientificCalc))
	fmt.Println(getSub(scientificCalc))
	fmt.Println(getMul(scientificCalc))
	fmt.Println(getDiv(scientificCalc))
}
