package main

import (
	"fmt"
	"strings"
)

func main() {

	/* Sprintf is a function in the fmt package in Go that formats and returns a string. It works similarly to Printf,
	  but instead of printing 
	  the formatted string to an output stream (like standard output), it returns the formatted string as a new string.*/
	fmt.Println("Sprintf Example : ")
    name := "Alice"
    age := 30

    // Sprintf formats the string and returns it
    formattedString := fmt.Sprintf("Hello, %s! You are %d years old.\n", name, age)

    // Now you can use the formatted string as needed
    fmt.Println(formattedString)


    // What is String Package?
	/*
	The strings package in Go provides functions for manipulating strings. 
	It is part of the standard library and offers a variety of useful functions for working with strings.
	*/
	// strings.Contains
	fmt.Println("Strings Package Functions Examples: ")
	str := "Hello, world!"
	contains := strings.Contains(str, "world")
	fmt.Println("Contains 'world':", contains) // true

	// strings.ToLower and strings.ToUpper
	str = "GoLang"
	lower := strings.ToLower(str)
	upper := strings.ToUpper(str)
	fmt.Println("ToLower:", lower) // golang
	fmt.Println("ToUpper:", upper) // GOLANG

	// strings.Split
	str = "apple,orange,banana"
	fruits := strings.Split(str, ",")
	fmt.Println("Split:", fruits) // [apple orange banana]

	// strings.Join
	fruitsJoined := strings.Join(fruits, " & ")
	fmt.Println("Join with '&':", fruitsJoined) // apple & orange & banana

	// strings.HasPrefix and strings.HasSuffix
	str = "Hello, world!"
	fmt.Println("Length of String 'Hello, world!' :" ,len(str))
	startsWith := strings.HasPrefix(str, "Hello")
	endsWith := strings.HasSuffix(str, "world!")
	fmt.Println("Starts with 'Hello':", startsWith) // true
	fmt.Println("Ends with 'world!':", endsWith)    // true

	str = "This is a sample sentence. This sentence has some repetition."
	replaced := strings.Replace(str, "sentence", "phrase", -1) //-1: Replace all occurrences of the target substring.
	fmt.Println("After Replace:", replaced)

	// strings.Index
	index := strings.Index(str, "sample")
	fmt.Println("Index of 'sample':", index)

	// strings.TrimSpace
	whitespaceStr := "   This is a string with leading and trailing spaces.   "
	trimmedStr := strings.TrimSpace(whitespaceStr)
	fmt.Println("Original String:", whitespaceStr)
	fmt.Println("Trimmed String:", trimmedStr)



	// String Replacer
	// The strings package in Go provides a Replace function that replaces all occurrences of a substring with another substring.
    fmt.Println("\nString Replacer Example:")
	string1 := "A word"

	replacer := strings.NewReplacer("A","Another")
	result := replacer.Replace(string1)
	fmt.Println(result)
}
