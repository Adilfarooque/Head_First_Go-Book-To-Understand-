package makingthegrade

import (
	"bufio"
	"fmt"
	"log"
	"os"
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
	fmt.Println(input)
}
