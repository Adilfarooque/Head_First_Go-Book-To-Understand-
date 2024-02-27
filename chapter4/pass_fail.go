package chapter4

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func GetFloat() (float64, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}

	input = strings.TrimSpace(input)
	number, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, err
	}
	return number, nil
}

func Grade() {
	fmt.Print("Enter a grade: ")
	grade, err := GetFloat()
	if err != nil {
		log.Fatal(err)
	}

	var status string
	if grade > 60 {
		status = "Passed"
	} else {
		status = "Failed"
	}
	fmt.Println("A grade of ", grade, "is", status)
}
