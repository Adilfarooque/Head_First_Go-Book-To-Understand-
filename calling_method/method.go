package callingmethod

import (
	"fmt"
	"time"
)

func CallingMehod() {
	var now time.Time = time.Now()
	var year int = now.Year()
	fmt.Println(year)
}
