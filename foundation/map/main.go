package main

import (
	"fmt"
	"sort"
)

// What is a map?
// A map is a built-in type in Go that is used to store key-value pairs. Maps are similar to Python dictionaries or Java HashMaps. They are essentially a reference to a hash table, which is a data structure that associates keys with values and allows for efficient retrieval of values based on their associated keys.

func main() {
	salary := map[string]int{
		"John":  2500,
		"Jane":  3200,
		"Jack":  5000,
		"Jill":  9000,
		"Joe":   0,
		"James": 5500,
	}

	// delete "Joe"
	delete(salary, "Joe")

	// Set new salary for "Jane" 3200 -> 4000
	salary["Jane"] = 4000
	fmt.Printf("The new salary of Jane is %d\n", salary["Jane"])

	// Check if "Jack" is present in the map
	if _, found := salary["Jack"]; found {
		fmt.Println("Jack is present")
	} else {
		fmt.Println("Jack is not present")
	}

	// A filter to show the salary > 3000
	for key, value := range salary {
		if value > 3000 {
			fmt.Printf(" # The salary of %s is %d\n", key, value)
		} else {
			fmt.Printf("%s is not eligible for salary hike\n", key)
		}
	}

	// Scroll through the map with name and salary
	for name, salary := range salary {
		fmt.Printf("The salary of %s is %d\n", name, salary)
	}

	// Scroll through the map ignoring the name with a blank identifier
	for _, salary := range salary {
		fmt.Printf("The salary is %d\n", salary)
	}

	// Make a sort for the map to take the highest salary
	// Extract the keys into a slice
	values := make([]int, 0, len(salary)) // Make is using to create a slice
	for _, value := range salary {
		values = append(values, value)
	}

	// Sort the keys
	sort.Sort(sort.Reverse(sort.IntSlice(values)))

	// Access the values in the original map using the sorted keys
	for _, values := range values {
		fmt.Printf("The highest salary is %d\n", values)
	}

}
