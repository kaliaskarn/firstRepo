package main

import (
	"fmt"
)

// exercise represents a simple struct with two fields.
type exercise struct {
	first  string
	second int
}

// methoding represents another struct with two string fields.
type methoding struct {
	killer    string
	assistant string
}

// String implements the Stringer interface for methoding.
func (me methoding) String() string {
	return fmt.Sprintf("The incident involved %s and their assistant %s.", me.killer, me.assistant)
}

// execution prints details about the exercise.
func (e exercise) execution() {
	fmt.Printf("All my friends are because of %s, there were a total of %d affected.\n", e.first, e.second)
}

// Executor is an interface that requires an execution method.
type Executor interface {
	execution()
}

// executingInterface calls the execution method of any Executor.
func executingInterface(exec Executor) {
	exec.execution()
}

func main() {
	e12 := exercise{
		first:  "Obama",
		second: 69,
	}

	me1 := methoding{
		killer:    "Andrew Tate",
		assistant: "Black",
	}

	executingInterface(e12)
	fmt.Println(me1)
}
