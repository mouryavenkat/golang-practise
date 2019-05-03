package main

import "fmt"

type functiontype = func(a int) int

// This interface.(type) can be only used with switch statement.
func handleInterfaceAccToType(assertionvar interface{}) {
	switch value := assertionvar.(type) {
	case int:
		fmt.Println("integer ")
	case functiontype:
		fmt.Printf("%T", value)
	}
}

// We can assign anything to a empty interface
func main() {
	var assertionvar interface{} = func(a int) int { return a }

	_, ok := assertionvar.(functiontype)

	fmt.Println(ok)

	handleInterfaceAccToType(assertionvar)
}
