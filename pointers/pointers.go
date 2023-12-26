package pointers

import "fmt"

func Pointers() {

	var myInt int
	fmt.Println(&myInt)

	var myFloat float64
	fmt.Println(&myFloat)

	var myBool bool
	fmt.Println(&myBool)


	amount := 6
	fmt.Println(amount)
	fmt.Println(&amount)

}
