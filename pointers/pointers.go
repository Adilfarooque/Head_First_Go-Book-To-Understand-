package pointers

import (
	"fmt"
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

	/*var myInt int
	fmt.Println(reflect.TypeOf(&myInt))

	var myFloat float64
	fmt.Println(reflect.TypeOf(&myFloat))

	var myBool bool
	fmt.Println(reflect.TypeOf(&myBool))

	*/
	/*
		myInt := 4
		myIntPointer := &myInt
		fmt.Println(myIntPointer)
		fmt.Println(*myIntPointer)

		myFloat := 98.6
		myFloatPointer := &myFloat
		fmt.Println(myFloatPointer)
		fmt.Println(*myFloatPointer)

		*myIntPointer = 19
		fmt.Println(*myIntPointer)
		fmt.Println(myInt)
	*/
	// var myFloatPointer *float64 = createPointer()
	// fmt.Println(*myFloatPointer)

	amount:= 6
	double(&amount)
	fmt.Println(amount)
}

func double(number *int){
	*number *= 2
}

// func createPointer() *float64 {
// 	var myFloat = 98.5
// 	//return a pointer of the specified type
// 	return &myFloat
// }
