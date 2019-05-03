package main

import "fmt"

// Any object assigned to interface, has to implement the methods defined in the interface
type testinterface interface {
	area() int
	perimeter() int
}

// Type implements interface by implementing its methods
type square struct {
	sidelength int
}
type rectangle struct {
	base   int
	height int
}

func (sq square) area() int {
	return sq.sidelength * sq.sidelength
}
func (rect rectangle) area() int {
	return (rect.base * rect.height) / 2
}
func (sq square) perimeter() int {
	return 0
}
func (rect rectangle) perimeter() int {
	return 0
}
func main() {
	var interface1, interface2 testinterface
	sq := &square{4}
	rect := &rectangle{4, 8}
	interface1 = sq // We canr assign square struct to testinterface struct unless we implement all the methods.
	//To know more abt interface , type (%T) and value (%v)
	fmt.Printf("%T %v", interface1, interface1)
	interface2 = rect
	fmt.Printf("%T %v", interface2, interface2)
	fmt.Println(interface1.area())
	fmt.Println(interface2.area())
}
