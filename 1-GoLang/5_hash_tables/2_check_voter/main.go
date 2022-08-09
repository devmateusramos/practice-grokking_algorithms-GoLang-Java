package main

import "fmt"

func main() {
	voted = make(map[string]bool)
	checkVoter("Tom")
	checkVoter("Mike")
	checkVoter("Tom")
}

var voted map[string]bool

func checkVoter(name string) {
	if voted[name] {
		fmt.Println("kick them out!")
	} else {
		voted[name] = true
		fmt.Println("let them vote!")
	}
}
