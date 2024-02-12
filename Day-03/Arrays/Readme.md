###Arrays in Go:

An array in Go is a numbered sequence of elements of a specific length and type.

###Characteristics of Go Arrays
<b>Fixed Size:</b> Once declared, the size of the array cannot be changed.
<b>Value Type:</b> Arrays in Go are value types, not reference types. When you assign or pass an array to a new variable, it copies the entire array.
<b>Zero Value:</b> Uninitialized elements of an array are set to the zero value of the array's type

##Different Ways to Declare Array in Go
###Declaration Without Initialization:
```
var arr [5]int
```
This declares an array arr of 5 integers. All elements are automatically initialized to the zero value of the type, which is 0 for integers.

###Declaration With Initialization:
```
var arr [5]int = [5]int{1, 2, 3, 4, 5}
```
This declares an array arr of 5 integers and initializes it with the values 1, 2, 3, 4, 5.

###Short Hand Declaration:
```
arr := [5]int{1, 2, 3, 4, 5}
```
This is a shorthand syntax for declaring and initializing an array. It infers the type based on the provided values.

###Initializing With Specific Elements:
```
arr := [5]int{1: 10, 3: 30}
```
This initializes an array of 5 integers where the second element (index 1) is 10 and the fourth element (index 3) is 30. Other elements are set to 0.
If an array where most elements are set to their zero value, it’s called sparse array

###Using the ... Operator to Infer Length
```
arr := [...]int{1, 2, 3, 4, 5}
```
Here, the length of the array is inferred from the number of initializer elements. The array arr will be of length 5.