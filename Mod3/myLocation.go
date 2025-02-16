package main

import (
	"fmt"
)

type Location struct {
	name    string
	country string
}

func New(name string, country string) Location {
	loc := Location{name, country}
	return loc
}

func myLocation(loc Location) {
	fmt.Println("Hi my name is " + loc.name + " and I live in " + loc.country + ".")
}

func main() {
	var loc1 Location = New("Steve", "Portugal")
	var loc2 Location = New("Dave", "Paris")
	var loc3 Location = New("Gwen", "Wales")

	myLocation(loc1)
	myLocation(loc2)
	myLocation(loc3)
}
