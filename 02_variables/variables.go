package main

import "fmt"

func main() {
	var username string = "Raghul M"
	fmt.Printf("Hello %s Welcome", username)
	//Type  of username
	fmt.Printf("\nType of username is %T \n", username)

	var isstring = true
	fmt.Println(isstring)
	//Type  
	fmt.Printf("Type of isstring is %T \n", isstring)

	var smallval uint8 = 255
	fmt.Println(smallval)
	//Type  
	fmt.Printf("Type of smallVal is %T", smallval)
}