package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type User struct {
	First string
	Age   int
}

func main() {
	u1 := User{
		First: "james",
		Age:   32,
	}

	u2 := User{
		First: "MoneyPenny",
		Age:   27,
	}
	u3 := User{
		First: "M",
		Age:   54,
	}
	users := []User{u1, u2, u3}

	fmt.Println(users)

	bs, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(bs))
	json.NewEncoder(os.Stdout).Encode(users)

}
