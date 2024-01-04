package pointers

import "fmt"

func PointerPool() {
	var myInt int

	var myIntPointer *int

	myInt = 42

	myIntPointer = &myInt

	fmt.Println(*myIntPointer)

}
