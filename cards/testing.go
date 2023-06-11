package main

import (
	"fmt"
)

type randomString []string

func rString() randomString {
	strings := randomString{}
	someRandomStrings := []string{"Spades", "Diamonds", "Hearts", "Club"}
	for _, v := range someRandomStrings {
		strings = append(strings, v)
		// fmt.Println(strings)
	}
	return strings
}

func (s randomString) printer() {
	for i, v := range s {
		fmt.Println(i, v)
	}
}

func main() {
	strings := rString()
	strings.printer()
}
