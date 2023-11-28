package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
cv1 := "5000"
cv3 := 1430

// What is strconv?
/* strconv is the package that provides functions for converting string to int, int to string,
 string to bool etc.*/

//Other Conversions in Strconv
/*

1. strconv.ParseFloat: Converts a string to a floating-point number (float64).
2. strconv.ParseInt and strconv.ParseUint: ParseInt parses a string into an integer of the specified bit size, 
   and ParseUint parses a string into an unsigned integer of the specified bit size.
3. strconv.FormatFloat: Converts a floating-point number (float64) to a string.
4. strconv.Itoa and strconv.Atoi: Itoa converts an integer to a string, and Atoi converts a string to an integer.


*/

cv2,_ := strconv.Atoi(cv1)  // Atoi - Ascii to Integer
fmt.Println("Atoi - Ascii to Integer")
fmt.Println(cv1, reflect.TypeOf(cv1))
fmt.Println(cv2, reflect.TypeOf(cv2))

cv4 := strconv.Itoa(cv3)  // Itoa - Integer to Asci
fmt.Println("Itoa - Integer to Asci")
fmt.Println(cv3, reflect.TypeOf(cv3))
fmt.Println(cv4, reflect.TypeOf(cv4))
fmt.Println("Using %T for Type")
fmt.Printf("%s %T\n", cv4, cv4) // By default %T returns Type of a variable 

}