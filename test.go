package main

import (
	"fmt"
	"strings"
)

func caps(name string) string {
	return strings.ToUpper(name)
}
func main() {
	var name string
	fmt.Println("Enter your name:")
	fmt.Scanln(&name)
	name_caps := caps(name)
	fmt.Print("Hello " + name_caps)
}
