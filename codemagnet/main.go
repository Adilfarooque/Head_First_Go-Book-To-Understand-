package main

import "fmt"

func main() {
	var originalCount int
	var eatenCout int
	originalCount = 10
	eatenCout = 4
	fmt.Println("I started with", originalCount, "apples.")
	fmt.Println("Some jerk ate", eatenCout, "apples.")
	fmt.Println("There are", originalCount-eatenCout, "apples left.")
}
