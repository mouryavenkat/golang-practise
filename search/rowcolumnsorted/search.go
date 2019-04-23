package main

import "fmt"

//SearchElement ... Searches element in a row and column sorted matrix and return bool tur/fasle
func SearchElement(element int) bool {
	arr := [4][4]int{{10, 20, 30, 40},
		{15, 25, 35, 45},
		{27, 29, 37, 48},
		{32, 33, 39, 50}}
	fmt.Println(arr)
	i, j := 0, 3
	for i >= 0 && i <= 3 && j >= 0 && j <= 3 {
		if arr[i][j] > element {
			fmt.Println(i, j)
			j--
		} else if arr[i][j] < element {
			fmt.Println(i, j)
			i++
		} else if arr[i][j] == element {
			return true
		}
	}
	return false
}
func main() {
	fmt.Println(SearchElement(32))
}
