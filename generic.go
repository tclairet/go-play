package main

import (
	"fmt"
)

/*
  It should work on https://go2goplay.golang.org/
*/

func ToPointer[T any](objects []T) []*T{
	var pointers []*T
	for i := 0; i < len(objects); i++ {
		pointers = append(pointers, &objects[i])

	}
	return pointers
}

func Print[T any](s []T) {
	for _, v := range s {
		fmt.Print(v)
		fmt.Print(" ")
	}
	fmt.Print("\n")
}

func PrintValues[T any](objects []*T) {
	for _, object := range objects {
		fmt.Print(*object)
		fmt.Print(" ")
	}
	fmt.Print("\n")
}


func main() {
	values := []int{0, 1}
	pointers := ToPointer(values)
	Print(pointers)

	Print(values)
	Print(pointers)
	PrintValues(pointers)

	*pointers[0] = 42

	Print(values)
	Print(pointers)
	PrintValues(pointers)
}

// this just a dummy change

// this just a second dummy change
