package main

import "fmt"

func main() {
	fmt.Println("Array:")
	var arr [5]int = [5]int{1, 2, 3, 4, 5} // array of five integers
	var arr1 [5]int = [5]int{1, 2, 3, 4} // It returns 5th value as 0 
	fmt.Println("arr=", arr)
	fmt.Println("arr1=", arr1)
}