package main

import (
	"fmt"
	"strconv"
)

func wait(what string, how int) (string, interface{}) {
	if what != "inst" {
		if how < 0 {
			how = how * -1
		}
	}
	if what == "inst" {
		fmt.Printf("%s: %d instructions \n", Oif(how > 0, "skipped", "move back"), Oif(how > 0, how, how*-1))
		return "move", how
	}
	fmt.Printf("wait: %d %s \n", how, what)
	return "wait", strconv.Itoa(how) + what
}

func motor(what int, how int) {
	if how > 100 {
		how = 100
	}
	if how < -100 {
		how = -100
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
	for _, e := range *varRef {
		if e.Name == what {
			return true, &e
		}
	}
	return false, nil
}
func Oif(s bool, t, f interface{}) interface{} {
	if s {
		return t
	}
	return f
}
