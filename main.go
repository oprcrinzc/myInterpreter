package main

import (
	"fmt"
	"os"
)

var fileExtension string = ".rbc"
var instructionSet [][]byte
var commands []string = []string{"motor", "servo", "set"}
var variables []Variable

func check(e error) {
	if e != nil {
		fmt.Println("Error!: " + e.Error())
		panic(e)
	}
}

func isIn[T comparable](what T, set []T) bool {
	for _, e := range set {
		if e == what {
			return true
		}
	}
	return false
}

func LoadInstructionSet(data []byte, where *[][]byte) {
	var instruction []byte
	for i, e := range data {
		if e == 32 || (e == 10) {
			// instruction = instruction[len(instruction)-1:]
			if len(instruction) > 0 {
				*where = append(*where, instruction)
			}
			instruction = nil
			continue
		}
		instruction = append(instruction, e)
		if i == len(data)-1 {
			if len(instruction) > 0 {
				*where = append(*where, instruction)
			}
		}
	}
}

func Execute(ins [][]byte) {
	var nextIndex *int32
	for i, e := range ins {
		if nextIndex != nil && i < int(*nextIndex) {
			fmt.Println(fmt.Sprintf("wait fot %d, %d", int(*nextIndex), i))
			continue
		} else {
			nextIndex = nil
		}
		if nextIndex == nil && isIn(string(e), commands) {
			nextIndex = nil
			command := string(e)
			if command == "motor" {
				// motor(, int(ins[i+2]))
				fmt.Print(string(ins[i+1]))
				fmt.Print(" ")
				fmt.Println(string(ins[i+2]))
				var a int32 = int32(i + 3)
				nextIndex = &a
				fmt.Println(*nextIndex)

			}
		}
		// fmt.Print(i)
		// fmt.Println(string(e))
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
	// fmt.Println(string(data))
	// fmt.Println(data)
	LoadInstructionSet(data, &instructionSet)
	fmt.Println(instructionSet)
	Execute(instructionSet)
}
