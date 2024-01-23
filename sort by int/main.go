package main

import (
	"fmt"
	"sort"
)

// person struct represents a person with a name and age.
type person struct {
	First string
	Age   int
}

// ByAge is a type that is essentially a slice of person structs.
type ByAge []person

// Len is part of sort.Interface and returns the length of the collection.
// It is crucial for the sorting algorithm to understand the size of the collection.
func (a ByAge) Len() int { return len(a) }

// Swap is part of sort.Interface. It swaps the elements at indices i and j.
// This method is used to rearrange elements during the sorting process.
func (a ByAge) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

// Less is part of sort.Interface. It defines the order of elements for sorting.
// In this case, it returns true if the person at index i should come before the person at index j (sorted by age).
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

func main() {
	// Creating a slice of person structs.
	p1 := person{"James", 32}
	p2 := person{"Moneypenny", 28}
	p3 := person{"Q", 64}
	p4 := person{"M", 56}

	people := []person{p1, p2, p3, p4}

	// Printing the slice before sorting.
	fmt.Println("Before sorting:", people)

	// sort.Sort sorts the collection.
	// By implementing sort.Interface (Len, Swap, Less) on ByAge,
	// we define how the slice of person structs should be sorted.
	// This is crucial because without this implementation, sort.Sort wouldn't
	// know how to sort the slice as it relies on these methods to make sorting decisions.
	sort.Sort(ByAge(people))

	// Printing the slice after sorting.
	fmt.Println("After sorting:", people)
}
