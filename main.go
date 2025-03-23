package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")

	var a = 15
	fmt.Printf("The value of a=%d\n", a)

	const b = 5
	fmt.Println("a=", a)

	age := 20

	if age >= 18 {
		fmt.Println("You are an adult.")
	} else {
		fmt.Println("You are a minor.")
	}
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		if i%2 == 0 {
			fmt.Printf("The number is Even\n")
		} else {
			fmt.Println("The number is Odd\n")
		}
	}
}
