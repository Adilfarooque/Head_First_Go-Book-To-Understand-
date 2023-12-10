package makingthegrade

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func PassOrFail() {
	//Prompt the user to enter the grade
	fmt.Print("Enter a grade: ")
	//Set up a "buffered reader" that gets text from the keyword.
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	//if "errors not nil..."
	if err != nil {
		//Report the error and stop the program.
		log.Fatal(err)
	}
	//Trim the newline charachter from the inpout string.
	input = strings.TrimSpace(input)
	//Convert the string to a float64 value.
	grade, err := strconv.ParseFloat(input, 64)
	//Just as with ReadString report any error when converting.
	if err != nil {
		log.Fatal(err)
	}
	/*Declare the "status" variable here, so
	it's in scope for the rest of the function*/
	var status string
	//Compare to the float64 in "grade", not the string in input.
	if grade >= 60 {
		status = "Passed"
	} else {
		status = "Failed"
	}
	//The pass or fail status
	fmt.Println("A greade of", grade, "is", status)

}
