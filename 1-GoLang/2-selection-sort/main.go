package main

import "fmt"

func main() {
	arrayToSort := []int{3, 6, 20, 2, 7}
	fmt.Println(selectionSort(arrayToSort))
}

// Finds the smallest value in an array
func findSmallest(arr []int) int {
	// Stores the Smallest value
	smallest := arr[0]
	// Stores the index of the smallest value
	smallest_index := 0
	for i := 1; i < len(arr); i++ {
		if arr[i] < smallest {
			smallest = arr[i]
			smallest_index = i
		}
	}
	return smallest_index
}

// Sort array
func selectionSort(arr []int) []int {
	size := len(arr)
	newArr := make([]int, size)
	for i := 0; i < size; i++ {
		smallest := findSmallest(arr)
		newArr[i] = arr[smallest]
		arr = append(arr[:smallest], arr[smallest+1:]...)
	}
	return newArr
}
