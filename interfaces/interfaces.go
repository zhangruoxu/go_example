package main

import (
	"fmt"
)

type Ball interface {
	Play()
}

type Football struct {
	owner string
}

func (f Football) Play() {
	fmt.Printf("%v plays football.\n", f.owner)
}

// ============================================
// Assign a variable to an interface will copy that variable.
// Assign a pointer to an interface can solve this issue.

func assignToInterface() {
	football := Football{owner: "Yifei"}
	var ball Ball = football
	ball.Play()
	football.owner = "Yifei Zhang"
	football.Play()
	ball.Play()

	var ballPtr Ball = &football
	football.Play()
	ballPtr.Play()
}

// ============================================

func modifyInterfaceVariable() {
	football := Football{owner: "Yifei"}
	var ball Ball = football
	ball.Play()
	football.Play()
	// The following two statements won't compile.
	// You cannot take the address of the struct referring by an interface.
	// ball.(Football).owner = "Yifei Zhang"
	// fmt.Printf("Address: %p", &ball.(Football).owner)
}

// ============================================

func modifyInterfacePointer() {
	football := Football{owner: "Yifei"}
	var ball Ball = &football
	ball.Play()
	football.Play()
	fmt.Printf("Obtaining the address of owner from the struct: %p\n", &(football.owner))
	fmt.Printf("Obtaining the address of owner from interface: %p\n", &(ball.(*Football).owner))
	ball.(*Football).owner = "Yifei Zhang"
	ball.Play()
	football.Play()
}

func main() {
	// assignToInterface()
	// modifyInterfaceVariable()
	modifyInterfacePointer()
}
