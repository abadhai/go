package main

import "fmt"

func main() {
	age := 13

	if age >= 18 {
		fmt.Println("Person is adult")

	} else if age>=12{
		fmt.Println("Person is teenager")
	}else{
		fmt.Println("Person is minor")
	}

	var role ="admin"
	var hasPermissions = true

	if role == "admin" || hasPermissions{

		fmt.Println("Yes")
	}

	if role == "admin" && hasPermissions{

		fmt.Println("Yes")
	}
}