package main

import (
	"errors"
	"fmt"
)

func main() {
	// sum
	value, err := sum(10, 20)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(value)

	fmt.Println(sum(52, 5))

	// area of a triangle
	result, err := area(0, 0)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)

	// fibonacci sequence
	fmt.Println(fibonacci(2))

  // sum with alot of numbers
  
  fmt.Println("The sum of the values is: ", sumNumbers(1,23,41,1231,4123,1231,2))

}

// use the last value to return an error to verify if the function is ok, this is very common in go
func sum(a int, b int) (int, error) { // retornar um valor
	if a+b >= 50 {
		return a + b, errors.New("valor muito alto") // retornar um erro
	}
	return a + b, nil
}

// area of a triangle
func area(width, height int) (int, error) {
	// test if the width is valid
	if width <= 0 {
		return width * height / 2, errors.New("Width is invalid")
	}

	return width * height / 2, nil
}

// fibonacci sequence algorithm 
func fibonacci(n int) ([]int, error) {
	// test if the number/length is valid
	if n <= 1 {
		return nil, errors.New("Invalid number")
	}

	list := make([]int, n) // create a slice with n positions

	list = []int{0, 1} // Initialize the first two positions

	for i := 2; i < n; i++ { // loop to fill the slice
		list[i] = list[i-1] + list[i-2]
	}

	return list, nil // return the slice
}

// Work with alot of params if we dont know 
func sumNumbers(numbers ...int) int{
    total := 0

    for _, number := range numbers{
      total += number
    }

    return total  
}
