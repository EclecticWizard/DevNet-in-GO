package main

import (
	"fmt"
)

func myCity(city string) {
	fmt.Println("I live in " + city + ".")
}

func main() {
	myCity("Texas")
	myCity("Toronto")
	myCity("London")
}
