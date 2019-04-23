package main

import "fmt"

func bubblesort() [5]int {
	arr := [5]int{7, 10, 2, 3, 8}
	arrlen := len(arr)
	for i := 0; i < arrlen-1; i++ {
		for j := i; j < arrlen; j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr
}

// Sort in increasing order using bubblesort
func main() {
	fmt.Println(bubblesort())
}
