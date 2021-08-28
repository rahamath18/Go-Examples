// Go - Reading file into string

package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {

	content, err := ioutil.ReadFile("d:/java/golang/read/readme.txt")
	if err != nil {
		log.Fatal(err)
	}

	text := string(content)
	fmt.Println(text)
}
