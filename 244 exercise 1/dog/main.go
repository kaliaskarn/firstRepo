package dog

import "fmt"

// Years take in an input of human years and turns it into dog years, via multiplying it by 7
func Years(h int) string {
	return fmt.Sprintln("His age in dog years will be", h*7)

}
