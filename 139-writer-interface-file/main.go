package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

type person struct {
	first string
}

// just like in a previous example in another code, io.Writer makes all the other types that use Writer Interface, available to use
// with the writeOut function when executing it. and its using Writer from the IO function.
func (p person) writeOut(w io.Writer) {
	w.Write([]byte(p.first))

}

func main() {

	p := person{
		first: "Jenny",
	}

	f, err := os.Create("output.txt")
	if err != nil {
		log.Fatalf("error %s", err)
	}

	defer f.Close()

	var b bytes.Buffer

	p.writeOut(f)
	p.writeOut(&b)
	fmt.Print(b.String())
}
