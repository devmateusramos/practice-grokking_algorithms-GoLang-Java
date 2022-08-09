package main

import "fmt"

func main() {
	greet("Mateus")
}

func greet2(name string) {
	fmt.Println("How are you, " + name + "?")
}

func bye() {
	fmt.Println("ok bye!")
}

func greet(name string) {
	fmt.Println("Hello, " + name + "!")
	greet2(name)
	fmt.Println("getting ready to say bye...")
	bye()
}
