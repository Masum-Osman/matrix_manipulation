package main

/*
#cgo CXXFLAGS: -std=c++11
#include <stdlib.h>
extern void multiply_matrices(int* a, int* b, int* result, int n);
*/
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Println("Starting the matrix multiplication...")

	n := 3 // Size of the matrices (3x3)

	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9} // 3x3 matrix
	b := []int{9, 8, 7, 6, 5, 4, 3, 2, 1} // 3x3 matrix

	result := make([]int, n*n)
	C.multiply_matrices((*C.int)(unsafe.Pointer(&a[0])), (*C.int)(unsafe.Pointer(&b[0])), (*C.int)(unsafe.Pointer(&result[0])), C.int(n))

	fmt.Println("Result of matrix multiplication:")
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Printf("%d ", result[i*n+j])
		}
		fmt.Println()
	}
}
