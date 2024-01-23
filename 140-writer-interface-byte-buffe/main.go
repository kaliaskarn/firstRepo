package main

import (
	"bytes"
	"fmt"
)

func main() {
	b := bytes.NewBufferString("new buffer printed ")
	b.WriteString(", this one was added to the bufffer")
	fmt.Println(b.String())
	b.Reset()
	b.WriteString("this is printed after reset")
	fmt.Println(b.String())

}
