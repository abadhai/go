package main

import "fmt"

const age = 30

func main() {

	const name = "Arun"
	
	fmt.Println(name)
	fmt.Println(age)

	const(
		port =5000
		host = "localhost"
	)
	fmt.Println(port)
	fmt.Println(host)
}