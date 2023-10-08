package main

import "fmt"

// Dictionary "fmt" https://golang.org/pkg/fmt/
// %v represents the named value in its default format
// %T represents the type of the value
// %d expects value to be an INTEGER type of base 10
// %b expects value to be an integer type of base 2
// %s the bytes of STRING or slice
// %f expects value to have a float type

// With go we can create a type for a variable, so we can leave the program more intuitive and easy to read.

type WORD string

// variable declaration global scope
var (
	feature bool
	num     int
	word    string
	float   float64
	custom  WORD = "Hello World" // My custom type
) // Declare multiple variables to show default values

// Go is declarative language, a way to declare variables is use := operator to infer the type for the first time. So with this we can declare and assign automatically the type of variable and aumaticaly use "=" for update the value of the same variable.

func main() {
	// variable declaration local scope
	// first time
	feature := false
	// second time its not necessary to use ":=" operator
	feature = true

	people_name := "John Doe"
	people_age := 30

	fmt.Println(feature)
	fmt.Println(custom)
	fmt.Printf("Your name is %s and your age is %d\n", people_name, people_age)
	fmt.Printf("Your name is %v and your age is %v\n", people_name, people_age)

	// create a list in go
	list := [5]int{1, 2, 3, 4, 5}
	fmt.Println(len(list))
	fmt.Println(list[0:2])

	// create a matrix in go
	matrix := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Printf("This is a matrix %v\n", matrix)

	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			empty := [2][3]int{}
			empty[i][j] = i + j
			fmt.Printf("Num: %d\n", empty[i][j])
		}
	}

	// create a map in go
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	fmt.Printf("This is a map in go %v\n", m)

	// SLICE
	// A slice is a segment of an array. Like arrays slices are indexable and have a length. Unlike arrays this length is allowed to change. Here's an example of a slice:

}
