package gussgame

import (
	"fmt"
	"math/rand"
)

func Guss() {
	target := rand.Intn(100) + 1
	fmt.Println(target)
}
