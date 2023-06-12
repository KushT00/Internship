package main

import "fmt"

func main2() {
	var name string
	fmt.Println("Enter Your name")
	fmt.Scanln(&name)
	fmt.Printf("hello " + name)
}
