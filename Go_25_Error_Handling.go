// Go - Error Handling

package main

import (
	"errors"
	"fmt"
)

type error interface {
	Error() string
}

func main() {
	err := isPositiveNumber(-12)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Given number is +ve")
	}
}

func isPositiveNumber(num int) error {
	if num < 0 {
		return errors.New("Negative Number Exception")
	}
	return nil
}
