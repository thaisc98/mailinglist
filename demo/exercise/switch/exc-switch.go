//--Summary:
//  Create a program to display a classification based on age.
//
//--Requirements:
//* Use a `switch` statement to print the following:
//  - "newborn" when age is 0
//  - "toddler" when age is 1, 2, or 3
//  - "child" when age is 4 through 12
//  - "teenager" when age is 13 through 17
//  - "adult" when age is 18+

package main

import "fmt"

func main() {
	switch age := 0; {
	case age == 0: 
		fmt.Println("newborn")
	case age <= 3 && age > 0:
		fmt.Println("toddler")
	case age <= 12 && age > 3:
		fmt.Println("child")
	case age <= 17 && age > 12:
		fmt.Println("teenager")
	default:
		fmt.Println("Adult")
	}

}

