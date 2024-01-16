package main
import "fmt"

func main(){
	const pi = 3.14
    // Cannot Modify the constant
	// Pi = 400 // enbling this line will through error: cannot assign to Pi
	fmt.Println("PI = ",pi)
}