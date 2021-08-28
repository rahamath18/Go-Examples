// Go - Reading file line by line

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("d:/java/golang/read/readme.txt")
	if err != nil {
		fmt.Println("Error: ", err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
