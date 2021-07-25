package main

import (
	"fmt"
	"reflect"
)

type X int

// ============================================
// Note the different between type and kind.

func examineTypes() {
	var a int = 100
	var b X = 100

	aType := reflect.TypeOf(a)
	bType := reflect.TypeOf(b)
	fmt.Printf("Type %v, kind %v\n", aType, aType.Kind())
	fmt.Printf("Type %v, kind %v\n", bType, bType.Kind())
	fmt.Printf("Type of type %T\n", aType)
}

// ============================================
// Constructing types.
// Note that array with different lengh are different types, but they have the same kind.
// Note that multiple level of pointer type can also be recorded.
// Elem() returns the basic type of pointers, slices, arrays, etc.

func constructingTypes() {
	a := reflect.ArrayOf(10, reflect.TypeOf(int(0)))
	fmt.Printf("Type %v, kind %v\n", a, a.Kind())

	b := reflect.ArrayOf(20, reflect.TypeOf(int(0)))
	fmt.Printf("Type %v, kind %v\n", b, b.Kind())

	fmt.Printf("Is type equals? %v. Is kind equals? %v\n", a == b, a.Kind() == b.Kind())

	c := reflect.MapOf(reflect.TypeOf(""), reflect.TypeOf(0))
	fmt.Printf("Type %v, kind %v\n", c, c.Kind())

	x := 0
	ptr := &x
	ptrPtr := &ptr
	d := reflect.TypeOf(x)
	ptrType := reflect.TypeOf(ptr)
	ptrPtrType := reflect.TypeOf(ptrPtr)

	fmt.Printf("Type %v, kind %v\n", d, d.Kind())
	fmt.Printf("Pointer type %v, kind %v, element type %v\n", ptrType, ptrType.Kind(), ptrType.Elem())
	fmt.Printf("Pointer pointer type %v, kind %v, element type %v\n", ptrPtrType, ptrPtrType.Kind(), ptrPtrType.Elem())
}

// ============================================
// Use reflection to examine structs.

type user struct {
	name string
	age  int
}

type manager struct {
	user
	title string
}

func ExamineStructs() {
	var m manager
	t := reflect.TypeOf(m)
	fmt.Printf("Type %v, kind %v, number of fields %v\n", t, t.Kind(), t.NumField())

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		fmt.Printf("Field name %v, type %v, offset %v, is anonymous %v\n", f.Name, f.Type, f.Offset, f.Anonymous)

		if f.Anonymous {
			for j := 0; j < f.Type.NumField(); j++ {
				af := f.Type.Field(j)
				fmt.Printf("  Field name %v, type %v, offset %v, is anonymous %v\n", af.Name, af.Type, af.Offset, af.Anonymous)
			}
		}
	}

	fmt.Println()

	// Find an anonymous field by its type name.
	user, _ := t.FieldByName("user")
	fmt.Printf("Field name %v, type %v, offset %v, is anonymous %v\n", user.Name, user.Type, user.Offset, user.Anonymous)

	// Find a field by name, anonymous fields are also supported.
	name, _ := t.FieldByName("name")
	fmt.Printf("Field name %v, type %v, offset %v, is anonymous %v\n", name.Name, name.Type, name.Offset, name.Anonymous)

	// Find a field by its indice
	age := t.FieldByIndex([]int{0, 1})
	fmt.Printf("Field name %v, type %v, offset %v, is anonymous %v\n", age.Name, age.Type, age.Offset, age.Anonymous)
}

func main() {
	examineTypes()
	constructingTypes()
	ExamineStructs()
}
