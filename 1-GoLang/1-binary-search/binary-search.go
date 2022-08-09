package main

import "fmt"

func main() {
	fmt.Println(checkBin([]int{1, 2, 4, 7, 9, 30, 33}, 6))  // -1
	fmt.Println(checkBin([]int{1, 2, 4, 7, 9, 30, 33}, 30)) // 5
}
func checkBin(list []int, numberList int) int {
	low := 0
	high := len(list) - 1

	for low <= high {
		mid := (low + high) / 2
		if list[mid] == numberList {
			return mid
		}
		if list[mid] < numberList {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1
}
