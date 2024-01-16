package main

import "fmt"

type Person struct{
	Name string
	age  int
}

func main(){
	fmt.Println("Structs:")
	
	var P1 Person = Person{"Raghul",21}
    fmt.Println("Struct P1: ",P1)
	fmt.Println("Name: ",P1.Name)
	fmt.Println("Age : ",P1.age)
}