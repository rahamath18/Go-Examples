// Go - Reading in Console Input in Golang

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter your name: ")
	name, _ := reader.ReadString('\n')

	fmt.Print("Enter your email: ")
	email, _ := reader.ReadString('\n')

	fmt.Print("Enter your city: ")
	city, _ := reader.ReadString('\n')

	fmt.Println("\nYour Information is here\n..........................")
	fmt.Println("Your name: ", name)
	fmt.Println("Your email: ", email)
	fmt.Println("Your city: ", city)
}
