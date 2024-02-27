package chapter4

import (
	"fmt"
	"log"
)

func ToCelsius() {
	fmt.Print("Enter a temprature in Fahrenheit: ")
	fahrenheit, err := GetFloat()
	if err != nil {
		log.Fatal(err)
	}
	Celsius := (fahrenheit - 32) * 5 / 9
	fmt.Printf("%0.2f degrees Celsius \n", Celsius)
}
