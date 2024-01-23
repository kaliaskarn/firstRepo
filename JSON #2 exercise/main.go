package main

import (
	"encoding/json"
	"fmt"
)

//func MarshalIndent(v any, prefix, indent string) ([]byte, error)
//[{"First":"james","Age":32},{"First":"MoneyPenny","Age":27},{"First":"M","Age":54}]

func main() {
	var marshalled = []byte(`[{"First":"james","Age":32},{"First":"MoneyPenny","Age":27},{"First":"M","Age":54}]`)

	type People struct {
		First string
		Age   int
	}

	var peoples []People
	err := json.Unmarshal(marshalled, &peoples)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(peoples)

	for i, People := range peoples {
		fmt.Println("the number of the person is: #", i)
		fmt.Println("\t \t", People.First, People.Age)
	}
}
