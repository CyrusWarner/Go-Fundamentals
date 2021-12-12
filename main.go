package main

import "fmt"

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
}
