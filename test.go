package main

import (
	"fmt"
	"strings"
)

func main() {
	var name string
	fmt.Println("Enter your name:")
	fmt.Scanln(&name)

	capitalizedName := strings.ToUpper(name)
	fmt.Printf("Hello %s\n", capitalizedName)
}
