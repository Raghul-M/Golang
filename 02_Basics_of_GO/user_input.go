package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var userinput string
 fmt.Println("Hello Enter your name: ")
 fmt.Scanf("%s",&userinput)

 fmt.Printf("Username: %s",userinput)

 // Using Bufio

 // Why Bufio?
 /* In this example, bufio.NewReader(os.Stdin) creates a buffered reader that reads from the standard input.
  The ReadString('\n') function reads a line of input until it encounters a newline character ('\n'). 
  This allows the program to read the entire line of input, including spaces,
   until the user presses Enter.*/

 fmt.Print("\nEnter some sentence: ")

	// Create a new bufio.Reader connected to standard input
	reader := bufio.NewReader(os.Stdin)

	// ReadString('\n') reads until the first occurrence of '\n' in the input
	userInput, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	fmt.Printf("Entered: %s", userInput)

}

