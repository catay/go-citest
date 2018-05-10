package main

import (
	"fmt"
	"github.com/pkg/errors"
	"os"
)

func main() {

	// bla
	fmt.Println("This is a test to make circleci work")

	f, err := os.Open("/tmp/bla.txt")
	if err != nil {
		fmt.Println(errors.Wrap(err, "open failed"))
	} else {
		f.Close()
	}

}
