package main

import (
	"fmt"
)

func motor(what int, how int) {
	if how > 100 {
		how = 100
	}
	if how < 0 {
		how = 0
	}
	fmt.Printf("motor: %d = %d \n", what, how)
}

func servo(what int, how int) {
	fmt.Printf("servo: %d = %d \n", what, how)
}

func makeVar(varRef *[]Variable, what string, how interface{}) {
	*varRef = append(*varRef, Variable{Name: what, Value: how})
}
func check(e error) {
	if e != nil {
		fmt.Println("Error!: " + e.Error())
		panic(e)
	}
}
func isErr(e error) bool {
	return e != nil
}

func isIn[T comparable](what T, set []T) bool {
	for _, e := range set {
		if e == what {
			return true
		}
	}
	return false
}
func isInVar(varRef *[]Variable, what string) (bool, *Variable) {

}
