package main

import (
	"fmt"
	"os"
	"strconv"
)

var fileExtension string = ".rbc"
var instructionSet [][]byte
var commands []string = []string{"motor", "servo", "set"}
var variables []Variable

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
			// fmt.Printf("wait for %d, %d \n", int(*nextIndex), i)
			continue
		} else {
			nextIndex = nil
		}
		if isIn(string(e), commands) {
			command := string(e)
			if command == "set" {
				p1 := string(ins[i+1])
				p2, err := strconv.Atoi(string(ins[i+2]))
				if isErr(err) {
					p2 := string(ins[i+2])
					makeVar(&variables, p1, p2)
				} else {
					makeVar(&variables, p1, p2)
				}
				a := int32(i + 3)
				nextIndex = &a
				fmt.Println(variables)
			}
			if command == "motor" {
				var p1i, p2i int
				// var p1s, p2s string
				var p1e, p2e error
				p1i, p1e = strconv.Atoi(string(ins[i+1]))
				if isErr(p1e) {
					b, r := isInVar(&variables, string(ins[i+1]))
					if b {
						p1i = r.Value.(int)
					}
				}
				p2i, p2e = strconv.Atoi(string(ins[i+2]))
				if isErr(p2e) {
					b, r := isInVar(&variables, string(ins[i+2]))
					if b {
						p2i = r.Value.(int)
					}
				}
				motor(p1i, p2i)
				a := int32(i + 3)
				nextIndex = &a
			}
			if command == "servo" {
				p1, err := strconv.Atoi(string(ins[i+1]))
				check(err)
				p2, err := strconv.Atoi(string(ins[i+2]))
				check(err)
				servo(p1, p2)
				a := int32(i + 3)
				nextIndex = &a
			}
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
	// fmt.Println(string(data))
	// fmt.Println(data)
	LoadInstructionSet(data, &instructionSet)
	fmt.Println(instructionSet)
	Execute(instructionSet)
}
