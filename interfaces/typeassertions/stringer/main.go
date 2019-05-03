package main

import "fmt"

type testStringer struct {
	firstname string
	lastname  string
}

func (stringobj testStringer) String() string {
	return fmt.Sprintf("Firstname is %v, Lastname is %v, FullName is %v", stringobj.firstname, stringobj.lastname, stringobj.firstname+stringobj.lastname)
}
func main() {
	stringerobj := &testStringer{"mourya", "Venkat"}
	fmt.Printf("%v", stringerobj) // By default print according to fmt package convetntion. But if it requires custom printing implement the String function on the type

	// If String Method is implemented, it access that method else takes default. ( Its like __repr__ or __str__ in python)
}
