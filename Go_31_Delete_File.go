// Go - File Delete

package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	fileName := "d:/java/golang/write/writeme.txt"
	err := os.Remove(fileName)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(fileName, " is deleted!")
	}
}
