package main

import (
	"fmt"
	"os"
)

var fileExtension string = ".rbc"
var instructionSet [][]byte

func check(e error) {
	if e != nil {
		fmt.Println("Error!: " + e.Error())
		panic(e)
	}
}

func LoadInstructionSet(data []byte) {
	var instruction []byte
	for i, e := range data {
		if e == 32 || (e == 10) {
			// instruction = instruction[len(instruction)-1:]
			instructionSet = append(instructionSet, instruction)
			instruction = nil
			continue
		}
		instruction = append(instruction, e)
		if i == len(data)-1 {
			instructionSet = append(instructionSet, instruction)
		}
	}
}

func main() {
	args := os.Args
	if fileName := args[1]; fileName[len(fileName)-4:] != fileExtension {
		fmt.Println(fileName)
		fmt.Println("i prefer .rbc file !")
		return
	}

	data, err := os.ReadFile(args[1])
	check(err)
	fmt.Println(string(data))
	fmt.Println(data)
	LoadInstructionSet(data)
	fmt.Println(instructionSet)
}
