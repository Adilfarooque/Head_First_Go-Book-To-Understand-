package gussgame

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
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
		seconds := time.Now().UnixNano()
		rand.New(rand.NewSource(seconds))
		//Call rand.Intn to generated numbers should be different each time!
		target := rand.Intn(100) + 1
		fmt.Println("")
		fmt.Println("I've chosen a random number b/w 1 and 100.")
		fmt.Println("Can you guess it?")

		//Create a bufio Reader which lets us read keyboard input
		reader := bufio.NewReader(os.Stdin)
		//Ask a number
		fmt.Println("")
		for guesses := 0; guesses < 10; guesses++ {
			
		fmt.Println("You have",10-guesses ,"guesses left.")

		fmt.Print("Make a guess: ")
		//Read what the user types up until they press Enter.
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		//Remove the newline
		input = strings.TrimSpace(input)
		//convert the input string to an integer
		guess, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}

		if guess < target {
			fmt.Println("OOps. Your guess was LOW.")
		} else if guess > target {
			fmt.Println("OOps. Your guess was HIGH.")
		} else if guess == target {
			fmt.Println("Extream Brain Boy")
		}

		fmt.Printf("Our guss is:- %d", target)
	}
	fmt.Println("")
}
