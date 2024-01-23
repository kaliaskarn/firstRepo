package main

import "fmt"

type person struct {
	first string
}

func changeStruct(p person) person {
	p.first = "He will always love her, always"
	return p
}

func changeStructP(p *person) {
	p.first = "She loves him back, a lot"

}

func main() {
	p1 := person{first: "Nurik loves aknur"}
	p2 := person{first: "Aknur loves Nurlan"}

	fmt.Println(p1)
	fmt.Println(p2)

	fmt.Println(changeStruct(p1))
	fmt.Println(p1)
	changeStructP(&p2)
	fmt.Println(p2)

}
