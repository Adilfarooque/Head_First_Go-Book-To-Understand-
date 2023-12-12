package gussgame

import (
	"fmt"
	"math/rand"
	"time"
	/*
		Import Path       package name

		"archive"  			archive
		"math/rand"			rand
		"archive/zip" 		zip
	*/)

func Guss() {
	/*To get different random numbers, we need to pass a value to the rand.Seed function
	That will "seed" random number generator 
	*/
	//Get the current date and time, as an integer.
	seconds:=time.Now().UnixNano()
	rand.New(rand.NewSource(seconds))
	//Call rand.Intn to generated numbers should be different each time!
	target := rand.Intn(100) + 1
	fmt.Println("I've chosen a random number b/w 1 and 100.")
	fmt.Println("Can you guess it?")
	fmt.Println(target)

}
