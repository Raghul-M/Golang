package main

import
(
	"fmt"
)

var pl = fmt.Println
var Age int
func main() {
	pl("Enter Age:")
	fmt.Scanf("%d",&Age)
	if (Age >= 1) && (Age <= 17){
		pl("Minor")
	}else if (Age > 18) || (Age == 18){
		pl("Major ")
	}else if ( Age == 0) {
		pl("New Born")
	}else{
		pl("Something Wrong!!")
	}

	pl("Example for Not Equal to ")
    
	number1 := 10
    number2 := 20

    // != means "not equal to"
    if number1 != number2 {
        fmt.Printf("%d is not equal to %d\n", number1, number2)
    } else {
        fmt.Printf("%d is equal to %d\n", number1, number2)
    }

}
