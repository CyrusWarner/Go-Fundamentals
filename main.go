// Has notes for the following sections of course
// Variable declarations and primitive data types
// Pointer data types and & operator for pointing to a pointers memory address
// Constants and Iota Constants

package main

import "fmt"

//everytime iota keyword is used it increments iotas value by 1
const (
	first  = iota + 6
	second = 2 << iota // 2 times 2 1 time as iota is equal to 1 since it is used once
	third
)

const (
	fourth = iota
)

func main() {
	// Variable declarations and primitive data types.

	var i int      // This is the initialized variable
	i = 42         // now we assign the initialized variable to 42
	fmt.Println(i) // prints what is passed in

	var f = 3.14 // float32Bit
	fmt.Println(f)

	firstName := "Cyrus" // := implicit initialization syntax implies the dataType of firstName based on what is assigned to it
	fmt.Println(firstName)

	b := true // boolean data type
	fmt.Println(b)

	c := complex(3, 4) // complex data type enables us to do complex math
	fmt.Println(c)

	r, im := real(c), imag(c) // real and imag functions for getting the values from a complex data type

	fmt.Println(r, im)

	// Begin pointer data types examples (pointer data type is a variable that points to another location that holds the data we are interested in)

	var middleName *string  // The asterisk before the value string is our pointer
	fmt.Println(middleName) // Will output nil as that is what is provided when we ask for a pointer object that points to no location

	var lastName *string = new(string) // initializes the pointer object
	*lastName = "Warner"               // putting asterisk before a pointer dereferences the pointer and reaches through the pointer and grabs the data and brings it back to you
	fmt.Println(*lastName)             // dereferences and prints the last names value

	ptr := &lastName       // the & value is a address pointer and points to the location of the last name pointer variable
	fmt.Println(ptr, *ptr) // if you dereference the address pointer ptr you will get the lastName value Warner

	newFirstName := "Cyrus"
	fmt.Println(newFirstName)

	newPtr := &newFirstName
	fmt.Println(newPtr, *newPtr) // prints the pointers memory address and the dereferenced value from the pointer

	newFirstName = "Riana"
	fmt.Println(newPtr, *newPtr) // memory address will still be the same but the dereferenced value is changed

	// Constants
	const pi = 3.1415 // Has to be initialized when it is declared. Cannot assign a const variable to a function that returns a value.
	fmt.Println(pi)

	const numValue int = 3
	fmt.Println(numValue + 3)

	fmt.Println(float32(numValue) + 1.2) // Allows us to add numValue int to the float32Bit number 1.2

	// Iota constants
	// Iota constants follow the same rule as Go built in constants
	// Iota resets between different constant blocks

	// first prints 6 as iota is 0 + 6
	// second prints 4 as iota is now 1 and the expression 2 >> iota is equal to 2 times 2 one time\
	// third is 8 as the value is not declared so it will add iota to 6.
	// fourth will be 0 as iota is being used in seperate constant block
	// << >> bitshift operators
	fmt.Println(first, second, third, fourth)

	// Arrays
	// long way of creating an array
	var arr [3]int //declares an array with 3 integers
	arr[0] = 1     // array at index 0 is equal to 1
	arr[1] = 2     // array at index 1 is equal to 2
	arr[2] = 3     // array at index 2 is equal to 3
	fmt.Println(arr)
	fmt.Println(arr[1]) // arrays are bounded so do not try to access an index that doesnt exist in the array

	// short way of creating an array
	arr2 := [3]int{1, 2, 3}
	fmt.Println(arr2)

	// Slices
	arr3 := [3]int{1, 2, 3}

	// creates a slice of the arr3 varialbe and assigns it to the slice variable
	// points to the data that the variable arr3 is holding so when printing the values they will be identical.

	slice := arr3[:]

	arr3[1] = 42
	slice[2] = 27

	fmt.Println(arr3, slice)

	slice2 := []int{1, 2, 3} // Go automatically resizes the array for us
	fmt.Println(slice2)

	slice2 = append(slice2, 4, 42, 27) // adding element to the slice2 array with the append function
	fmt.Println(slice2)

	slice3 := slice2[1:]  // Creates a slice of the value at index 1 and every index afer index 1
	slice4 := slice2[:2]  // Creates a slice of every value before index 2 and not including the value at index 2
	slice5 := slice2[1:2] // Creates a slice of the value at index 1 and everything before index 2 and not including index 2

	fmt.Println(slice3, slice4, slice5)

}
