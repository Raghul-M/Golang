package main

import (

"reflect"
"fmt"
)

var pl = fmt.Println

func main() {

	/*
	In Go, the reflect package provides a set of methods for runtime reflection, allowing you to examine and
	 manipulate the structure and behavior of types at runtime. Reflection is a powerful but advanced feature, 
	and its use should be approached with caution as it can lead to less type-safe and less efficient code.
	*/

	var name = "Raghul"
	pl(reflect.TypeOf(25))
	pl(reflect.TypeOf(3.4))
	pl(reflect.TypeOf(true))
	pl(reflect.TypeOf("hello"))
	pl(reflect.TypeOf('😀'))

	// reflect ValueOf 
	pl(reflect.ValueOf(name))

}