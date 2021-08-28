// Go - New File Write

package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Create("d:/java/golang/write/writeme.txt")

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	defer file.Close()

	len, err := file.WriteString("Welcome to Go." +
		" This program demonstrates writing" +
		" operations to a file in Go lang.")
	file.WriteString("Welcome to Go." +
		" This program demonstrates writing" +
		" operations to a file in Go lang.")

	if err != nil {
		log.Fatalf("failed writing to file: %s", err)
	}
	fmt.Println(len)
}
