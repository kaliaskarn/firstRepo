package main

import (
	"fmt"
)

func main() {
	m := map[string]int{
		"James":      42,
		"Moneypenny": 32,
	}
	for k, v := range m {
		fmt.Printf("index: %v \t key: %v\n", k, v)

		age := m["James"]
		fmt.Println(age)

		age2 := m["Q"]
		fmt.Println(age2)
	}

	if s, ok := m["Q"]; ok {
		fmt.Printf("the value of Q is %v\n", s)
	} else {
		fmt.Println("Q wasnt found")
	}
}

/*func offset(tz string) int {
    if seconds, ok := timeZone[tz]; ok {
        return seconds
    }
    log.Println("unknown time zone:", tz)
    return 0
}
*/
