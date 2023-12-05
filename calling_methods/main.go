package main

import (
	"fmt"
	"time"
)

func main() {
	var now time.Time = time.Now()
	var Day_year int = now.YearDay()
	fmt.Println(Day_year)
}
