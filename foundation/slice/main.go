package main

import "fmt"

// In Go (Golang), a slice is a dynamic, flexible, and reference-like data structure that represents a portion of an underlying array. Slices are used to work with sequences of data, and they are a fundamental data structure in Go. Unlike arrays, slices do not have a fixed length; they can grow or shrink as needed.

// Dynamic Length: Slices can change in size dynamically. You don't need to specify their length when declaring them.

// Reference to Underlying Array: Slices are essentially references to a contiguous segment of an array. This means that if you modify the elements of a slice, it will also affect the underlying array, and any other slices that share the same underlying array.

// Syntax: Slices are defined by specifying a start index (inclusive) and an end index (exclusive) within square brackets when slicing an array or another slice. For example, mySlice := myArray[1:4] creates a new slice mySlice that contains elements from index 1 up to, but not including, index 4 of myArray.

// Capacity: Slices have both a length (the number of elements they contain) and a capacity (the maximum number of elements they can hold without resizing). You can use the len() and cap() functions to retrieve these values.
// Append Function: To add elements to a slice, you can use the append() function, which may reallocate the underlying array if the capacity is exceeded.

func main() {
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Printf("len=%d cap=%d %v\n", len(array), cap(array), array)
	fmt.Printf("len=%d cap=%d %v\n", len(array[:0]), cap(array[:0]), array[:0])
	fmt.Printf("len=%d cap=%d %v\n", len(array[:4]), cap(array[:4]), array[:4])
	fmt.Printf("len=%d cap=%d %v\n", len(array[2:]), cap(array[2:]), array[2:])

	array = append(array, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25)

	// fmt.Printf("len=%d cap=%d %v\n", len(array[:2]), cap(array[:2]), array[:2])
	fmt.Printf("len=%d cap=%d %v\n", len(array), cap(array), array)

}
