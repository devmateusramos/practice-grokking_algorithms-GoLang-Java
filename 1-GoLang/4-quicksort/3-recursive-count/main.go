package main

import "fmt"

func main() {
	fmt.Println(countItems([]int{0, 1, 2, 10, 3, 4, 5, 8}))
}

// conta o número de items na lista
func countItems(list []int) int {
	if len(list) == 0 {
		return 0
	}

	return 1 + countItems(list[1:])
}
