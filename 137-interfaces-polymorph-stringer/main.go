package main

import (
	"fmt"
	"strconv"
)

// Define a struct 'book' with a field 'title'.
type book struct {
	title string
}

// Implement the String method for 'book'.
// This method makes 'book' satisfy the Stringer interface.
// It defines how a 'book' should be represented as a string.
func (b book) String() string {
	// Return a formatted string representing the book.
	return fmt.Sprint("The title of the book is ", b.title)
}

// Define a new type 'count' as an alias for 'int'.
type count int

// Implement the String method for 'count'.
// This method makes 'count' satisfy the Stringer interface.
// It defines how a 'count' should be represented as a string.
func (c count) String() string {
	// Convert the 'count' value to a string and return it.
	return fmt.Sprint("This is the number ", strconv.Itoa(int(c)))
}

func main() {
	// Create an instance of 'book'.
	b := book{
		title: "West With The Night",
	}

	// Create an instance of 'count'.
	var n count = 42

	// Use fmt.Println to print the book.
	// fmt.Println automatically calls the String method of 'book'.
	// This is because 'book' implements the Stringer interface.
	// So, instead of manually calling a method to print 'book',
	// fmt.Println handles it, providing a cleaner and more idiomatic way to print.
	fmt.Println(b)

	// Similarly, use fmt.Println to print the count.
	// fmt.Println automatically calls the String method of 'count'.
	// This provides a consistent way to print different types that implement Stringer.
	fmt.Println(n)
}
