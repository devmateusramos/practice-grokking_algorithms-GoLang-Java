package main

import "fmt"

func main() {
	countdown(5)
}

func countdown(i int) {
	fmt.Println(i)
	if i <= 0 {
		return
	}

	countdown(i - 1)
}
