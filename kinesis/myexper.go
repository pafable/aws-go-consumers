package main

import (
	"fmt"
)

func main() {
	r := "sup dog"
	s := []byte(r)
	fmt.Printf(" Type: %T\nValue: %v\n", s, s)
}
