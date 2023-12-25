package double

import "fmt"

func Double() {
	amount := 11
	Db(amount) //pass an argument to the function
	fmt.Println(amount)
}

//Parameter is set to a copy of the argument
func Db(number int) {
	number *= 2 //Alter the copied value, not the original!
}
