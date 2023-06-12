package main

import "fmt"

func main() {
	var name string
	fmt.Println("Enter Your name")
	fmt.Scanln(&name)
	fmt.Printf("hello " + name)

}
