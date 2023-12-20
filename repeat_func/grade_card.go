package repeatfunc

import "fmt"

func status(grade float64) string {
	if grade > 60.0 {
		return "Failed"
	}
	return "Passed"
}

func Grade() {
	fmt.Println(status(60.1))
	fmt.Println(status(50))
}
