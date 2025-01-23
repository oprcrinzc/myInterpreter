package main

import (
	"fmt"
	"os"
)

var fileExtension string = ".rbc"

func main() {
	args := os.Args
	if fileName := args[1]; fileName[4:] != fileExtension {
		fmt.Println("i prefer .rbc file !")
		return
	}

	fmt.Println("ok")
}
