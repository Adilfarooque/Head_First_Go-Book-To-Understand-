package repeatfunc

import "fmt"

var MtPerLiter float64

//Declare that paintNeeded will return a floating point number.
func paintNeeded(width, height float64) float64 {
	area := width * height
	return area / 10.0 //Return area instead of printing it
}
func MetersPerLiter() {
	var amount, total float64                   //Declare variables to hold the amount for the current wall, as well as the total for all walls.
	amount = paintNeeded(4.2, 3.0)              //Call paintNeeded , and store the return value.
	fmt.Printf("%0.2f liters needed\n", amount) //Print the amount for this wall.
	total += amount                             //Add the amount for this wall to the total.

	amount = paintNeeded(5.2, 3.5)
	fmt.Printf("%0.2f liters needed\n", amount)
	total += amount

	fmt.Printf("TOTAL : %0.2f liters \n", total)
}
