package main

import (
	"fmt"
	"log"
)

// timeZone is a map that contains time zones as keys and their offset in seconds as values.
var timeZone = map[string]int{
	"UTC": 0,      // Coordinated Universal Time
	"EST": -18000, // Eastern Standard Time (UTC-5)
	"PST": -28800, // Pacific Standard Time (UTC-8)
	// Add more time zones as needed
}

// offset function takes a time zone string and returns its offset in seconds.
func offset(tz string) int {
	if seconds, ok := timeZone[tz]; ok {
		return seconds
	}
	log.Println("unknown time zone:", tz)
	return 0
}
