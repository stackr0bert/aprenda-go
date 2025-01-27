package main

import (
	"fmt"
)

func main() {
	array := [5]int{1, 2, 3, 4, 5} // numero finito de elementos
	slice := []int{1, 2, 3, 4, 5}  // numero não especificado de elementos
	fmt.Println(array)
	fmt.Println(slice)
	slice2 := append(slice, 6)
	fmt.Println(slice2)

	fmt.Println(slice[3])
	slice[3] = 54321
	fmt.Println(slice)
}
