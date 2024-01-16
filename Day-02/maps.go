package main

import "fmt"

func main(){
	fmt.Println("Maps:")
	
	 P1 := map[string]int{"Raghul":21,"Raja":22}
    fmt.Println("Map P1: ",P1)
	fmt.Println("Raghul Age: ",P1["Raghul"])
	fmt.Println("Raja Age: ",P1["Raja"])
	
}