package repeatfunc

import "fmt"

var MtPerLiter float64

func paintNeeded(width, height float64) float64 {
	area := width * height
	return area / MtPerLiter
}
func MetersPerLiter() {
	MtPerLiter = 10.0
	fmt.Printf("%2f", paintNeeded(4.2,3.0))
}
