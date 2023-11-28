package main

import (
	"fmt"
	"unicode/utf8"
)

//What is Rune ?
/*Rune is an alias for int32 and is equivalent to int32 in all ways. It is used, by convention, 
to distinguish character values from integer values.*/


func main() {
	rStr := "abcdefg"
	fmt.Println("Rune Count :", utf8.RuneCountInString(rStr))
	for i, runeVal := range rStr{
		fmt.Printf("%d : %#U : %c\n",i,runeVal,runeVal)
	}
}