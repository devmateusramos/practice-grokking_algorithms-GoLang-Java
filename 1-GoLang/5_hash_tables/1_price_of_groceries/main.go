package main

import "fmt"

func main() {
	book = make(map[string]float64)
	book["apple"] = 0.67
	book["milk"] = 1.49
	fmt.Println(book)
}

var book map[string]float64
