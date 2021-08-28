// Go - Local Variables Scope Rules & Global Variables Scope Rules
package main

import "fmt"

var global int = 5

func main() {
	fmt.Println(global) // output = ?
	var local1 int = 10
	fmt.Println(local1) // output = ?
	global = 15
	fmt.Println(global) // output = ?
}
