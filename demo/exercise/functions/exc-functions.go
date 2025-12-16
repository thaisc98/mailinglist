//--Summary:
//  Use functions to perform basic operations and print some
//  information to the terminal.
//
//--Requirements:
//* Add three numbers together using any combination of the existing functions.
//  * Print the result
//* Call every function at least once

package main

import (
	"fmt"
	"math/rand"
)

//* Write a function that accepts a person's name as a function
//  parameter and displays a greeting to that person.
func sayGreeting(person string){
	fmt.Println("Greetings ",person )
}

//* Write a function that returns any message, and call it from within
//  fmt.Println()
func anyMessage() string{
	return "Hello"
}

//* Write a function to add 3 numbers together, supplied as arguments, and
//  return the answer
func add(fst,scd,thrd int) int {
	return fst + scd + thrd 
} 


//* Write a function that returns any number
func randomNumber() int {
	return rand.Intn(10)
}

//* Write a function that returns any two numbers
func twoTwos () (int, int){
	return 2,2
}

func main() {
	sayGreeting("Thais")
	fmt.Println("Any message: ", anyMessage())
	
	answer := add(1,2,3)
	fmt.Println("the answer is ", answer)
	random := randomNumber()
	fmt.Println("the random number is ", random)
	value1, value2  := twoTwos()
	fmt.Println("the 2 random number is first", value1, "random second: ", value2)

	value := add(randomNumber(), value1, value2)
	fmt.Println("Random add", value)

}

