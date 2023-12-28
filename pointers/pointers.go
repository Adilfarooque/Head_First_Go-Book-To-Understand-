package pointers

import (
	"fmt"
	"reflect"
)

func Pointers() {
	/*
		var myInt int
		fmt.Println(&myInt)

		var myFloat float64
		fmt.Println(&myFloat)

		var myBool bool
		fmt.Println(&myBool)


		amount := 6
		fmt.Println(amount)
		fmt.Println(&amount)
	*/

	// Pointer types

	var myInt int
	fmt.Println(reflect.TypeOf(&myInt))

	var myFloat float64
	fmt.Println(reflect.TypeOf(&myFloat))

	var myBool bool
	fmt.Println(reflect.TypeOf(&myBool))
}
