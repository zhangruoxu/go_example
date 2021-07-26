package main

import (
	"fmt"
)

type user struct {
	name string
}

func (u user) sayHi(msg string) {
	fmt.Printf("Hi, %v. %v\n", u.name, msg)
}

func freeFunction(msg string) {
	fmt.Println(msg)
}

// ============================================
// Use the function name to refer a function.

func accessFunction() {
	function := freeFunction
	function("Hi, how are you?")
}

// ============================================
// Use type name to refer a methods.
// The first argument is the receiver object.

func accessMethod() {
	u := user{name: "Yifei"}
	mtd := user.sayHi
	mtdPtr := (*user).sayHi
	fmt.Printf("Method address %v, type %T\n", mtd, mtd)
	mtd(u, "Good to see you.")
	mtdPtr(&u, "Good to see you.")
}

// ============================================
// Method reference can be accessed from the instance by using the method name.
// Note that the returned value is actual a closure,
// which means the receiver instance has been partial applied.
//
// The type of the method reference obtained via the type is:
//
// func(method_owner, parameters).
//
// The type of the method reference obtained via the instance is:
//
// func(parameters), i.e., the receiver instance have already been applied.

func accessBoundMethod() {
	u := user{name: "Yifei"}
	mtd := u.sayHi
	fmt.Printf("Bound method %v, type %T\n", mtd, mtd)
	invokeClosure("Nice to meet you.", mtd)

	mtdPtr := (&u).sayHi
	fmt.Printf("Bound method %v, type %T\n", mtdPtr, mtdPtr)
	invokeClosure("Nice to meet you, too.", mtd)
}

func invokeClosure(msg string, mtd func(string)) {
	mtd(msg)
}

func main() {
	accessFunction()
	accessMethod()
	accessBoundMethod()
}
