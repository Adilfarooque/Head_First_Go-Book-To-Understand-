package repeatfunc

import "fmt"

var MtPerLiter float64

//Declare that paintNeeded will return a floating point number.
func paintNeeded(width, height float64) (float64, error) {
	if width < 0 {
		return 0, fmt.Errorf("a width of %0.2f is invalid", width)
	}
	if height < 0 {
		return 0, fmt.Errorf("a height of %0.2f is invalid", height)
	}

	area := width * height
	return area / 10.0, nil //Return area instead of printing it
}
func MetersPerLiter() {
	var amount, total float64            //Declare variables to hold the amount for the current wall, as well as the total for all walls.
	amount, err := paintNeeded(4.2, -3.0) //Call paintNeeded , and store the return value.
	fmt.Println(err)
	fmt.Printf("%0.2f liters needed\n", amount) //Print the amount for this wall.
	total += amount                             //Add the amount for this wall to the total.

	amount, err = paintNeeded(5.2, -3.5)
	fmt.Printf("%0.2f liters needed\n", amount)
	fmt.Println(err)
	total += amount

	fmt.Printf("TOTAL : %0.2f liters \n", total)
}
