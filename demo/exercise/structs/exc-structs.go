//--Summary:
//  Create a program to calculate the area and perimeter
//  of a rectangle.
//
//--Requirements:
//* Create a rectangle structure containing a length and width field
//* Using functions, calculate the area and perimeter of a rectangle,
//  - Print the results to the terminal
//  - The functions must use the rectangle structure as the function parameter
//* After performing the above requirements, double the size
//  of the existing rectangle and repeat the calculations
//  - Print the new results to the terminal
//
//--Notes:
//* The area of a rectangle is length*width
//* The perimeter of a rectangle is the sum of the lengths of all sides

package main

import "fmt"

type Rectangle struct {
	length int
	width int
}

func calculateAreaAndPerimeter (rect Rectangle) (int, int) {
	return rect.length * rect.width, 2 * (rect.length + rect.width)
}

func printInfo(rect Rectangle){
	area, perimeter := calculateAreaAndPerimeter(rect)
	fmt.Println("The area of the rectangle is: ", area)
	fmt.Println("The perimeter of the rectangle is: ", perimeter)
}

func main() {
	rectangle := Rectangle{length: 2, width: 2}
	printInfo(rectangle)

	rectangle.length *= 2
	rectangle.width *= 2

	printInfo(rectangle)

}


