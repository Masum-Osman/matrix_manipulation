// main.go
package main

/*
#cgo CXXFLAGS: -std=c++11
#include <stdlib.h>
extern int add(int a, int b);
*/
import "C"
import (
	"fmt"
)

func main() {
	a := 5
	b := 10

	// Call the C++ function
	result := C.add(C.int(a), C.int(b))

	// Print the result
	fmt.Printf("Result of adding %d and %d is %d\n", a, b, result)
}
