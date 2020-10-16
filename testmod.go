package testmod

import (
	"fmt"
)

//Hi returns a friendly greeting
func Hi(name string) string {
	return fmt.Sprintf("Hi, %s!", name)
}

//Bye says bye.
func Bye(name string) string {
	return fmt.Sprintf("Bye, %s!", name)
}
