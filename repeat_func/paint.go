package repeatfunc

import "fmt"

var MtPerLiter float64

func paintNeeded(width, height float64) float64 {
	area := width * height
	return area / 10.0
}
func MetersPerLiter() {
	var amount, total float64
	amount = paintNeeded(4.2, 3.0)
	fmt.Printf("%0.2f liters needed\n", amount)
	total += amount

	amount = paintNeeded(5.2, 3.5)
	fmt.Printf("%0.2f liters needed\n", amount)
	total += amount

	fmt.Printf("TOTAL : %0.2f liters \n", total)
}
