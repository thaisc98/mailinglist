//Summary:
//  Print basic information to the terminal using various variable
//  creation techniques. The information may be printed using any
//  formatting you like.
//
//Requirements:
//* Store your favorite color in a variable using the `var` keyword
//* Store your birth year and age (in years) in two variables using
//  compound assignment
//* Store your first & last initials in two variables using block assignment
//* Declare (but don't assign!) a variable for your age (in days),
//  then assign it on the next line by multiplying 365 with the age
// 	variable created earlier
//
//Notes:
//* Use fmt.Println() to print out information
//* Basic math operations are:
//    Subtraction    -
// 	  Addition       +
// 	  Multiplication *
// 	  Division       /

package main

import "fmt"

func main() {
	var myFavColor string = "Black"
	year, age := 1998, 27

	var (
		first = "T"
		last = "C"
	)

	var age2 int
	age2 = 365 * age

	fmt.Println("my fav color is ", myFavColor)
	fmt.Println("my year is ", year, " and my age is ", age)
	fmt.Println("My initial is ", first, "and my last is ", last)
	fmt.Println("My age in days is ", age2)


}

