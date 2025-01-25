package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

var fileExtension string = ".rbc"
var instructionSet [][]byte
var commands []string = []string{"motor", "servo", "set", "wait"}
var variables []Variable
var isEndProc bool

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

func Execute(inst [][]byte) {
	var nextIndex *int32
	var instSize int = len(inst)
	if instSize < 0 {
		return
	}
	var instMaxIndex = instSize - 1
	var instProc [][]byte
	fmt.Printf("inst size: %d\n", instSize)
	for !isEndProc {
		if nextIndex == nil {
			instProc = inst
		} else {
			instProc = inst[*nextIndex:]
			// fmt.Print(string(instProc[0]))
			// os.Exit(1)
		}
		for i, e := range inst {
			if nextIndex != nil && i < int(*nextIndex) {
				continue
			} else {
				nextIndex = nil
			}
			if isIn(string(e), commands) {
				command := string(e)
				if command == "wait" {
					var p2i int
					var p2e error
					p2i, p2e = strconv.Atoi(string(inst[i+2]))
					if isErr(p2e) {
						b, r := isInVar(&variables, string(inst[i+2]))
						if b {
							p2i = r.Value.(int)
						}
					}
					what, how := wait(string(inst[i+1]), p2i)
					// fmt.Printf("%s %d\n", what, how.(int))
					if what == "wait" {
						d, err := time.ParseDuration(how.(string))
						check(err)
						time.Sleep(d)
					}
					if what == "move" {
						f := 3*how.(int) + i + 3
						if f < 0 {
							f = instMaxIndex + f
						}
						fmt.Printf("f: %d, i: %d, max i: %d\n", f, i, instMaxIndex)
						if f > instMaxIndex {
							isEndProc = true
							break
						}
						m := int32(f)
						fmt.Printf("to index: %d\n", m)
						nextIndex = &m
						break
					}
				}
				if command == "set" {
					p1 := string(inst[i+1])
					p2, err := strconv.Atoi(string(inst[i+2]))
					if isErr(err) {
						p2 := string(inst[i+2])
						makeVar(&variables, p1, p2)
					} else {
						makeVar(&variables, p1, p2)
					}
				}
				if command == "motor" || command == "servo" {
					var p1i, p2i int
					var p1e, p2e error
					p1i, p1e = strconv.Atoi(string(inst[i+1]))
					if isErr(p1e) {
						b, r := isInVar(&variables, string(inst[i+1]))
						if b {
							p1i = r.Value.(int)
						}
					}
					p2i, p2e = strconv.Atoi(string(inst[i+2]))
					if isErr(p2e) {
						b, r := isInVar(&variables, string(inst[i+2]))
						if b {
							p2i = r.Value.(int)
						}
					}
					if command == "motor" {
						motor(p1i, p2i)
					} else {
						servo(p1i, p2i)
					}
				}
				a := int32(i + 3)
				nextIndex = &a
				// fmt.Println(variables)
			}
			// fmt.Println(len(instProc))
			if i >= len(instProc) {
				isEndProc = true
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
	LoadInstructionSet(data, &instructionSet)
	// fmt.Println(instructionSet)
	isEndProc = false
	Execute(instructionSet)
}
