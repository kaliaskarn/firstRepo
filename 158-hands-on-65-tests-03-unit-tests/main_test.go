package main

import (
	"log"
	"testing"
)

func TestRapist(t *testing.T) {
	weGot := Rapist("Petr")
	want := "Petr"
	if weGot != want {
		log.Fatalf("error  bitch")
	}

}
