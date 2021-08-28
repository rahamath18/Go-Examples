// Go - Reading in Console Input in Golang

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		fmt.Println(text)
		//if text == '$' {
		//goto Exit
		//}
	}
}
