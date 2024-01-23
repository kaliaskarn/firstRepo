package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	myPassword := "Maraqpoopiehole!"
	bs, err := bcrypt.GenerateFromPassword([]byte(myPassword), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(myPassword)
	fmt.Println(bs)

	passwordCheck := "Broski123123"
	err = bcrypt.CompareHashAndPassword(bs, []byte(passwordCheck))
	if err != nil {
		fmt.Println("its not the same")
		return
	}
	fmt.Println("they the same, password match")

}
