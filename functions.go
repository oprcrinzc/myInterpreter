package main

import "fmt"

func motor(what int, how int) {
	if how > 100 {
		how = 100
	}
	if how < 0 {
		how = 0
	}
	fmt.Println(fmt.Sprintf("motor: %d = %d", what, how))
}

func servo(what int, how int) {
	fmt.Println(fmt.Sprintf("servo: %d = %d", what, how))
}
