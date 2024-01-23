package main

import "fmt"

// dog struct represents a dog with a name.
type dog struct {
	name string
}

// walk is a method with a value receiver.
// It can be called on both dog and *dog.
func (d dog) walk() {
	fmt.Println("the dog that is walking, her name is ", d.name)
}

// run is a method with a pointer receiver.
// It modifies the dog's name and can only be called on *dog.
func (d dog) run() {
	d.name = "nigga dog"
	fmt.Println("the dog that is running right now, her name is ", d.name)
}

// execution interface requires both walk and run methods.
type execution interface {
	walk()
	run()
}

// runwalk takes an execution interface and calls both walk and run.
func runwalk(e execution) {
	e.walk()
	e.run()
}

func main() {
	// d1 is a dog value.
	d1 := dog{"garfield"}

	// d2 is a pointer to a dog.
	d2 := dog{"broski"}

	// runwalk is called with d2, which satisfies the execution interface.
	runwalk(d2)
	runwalk(d1)

	d2 = d2.changeName("Obama")
	runwalk(d2)

}

func (d dog) changeName(s string) dog {
	d.name = s
	return d
}
