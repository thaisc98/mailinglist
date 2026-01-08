//--Summary:
//  Create a calculator that can perform basic mathematical operations.
//
//--Requirements:
//* Mathematical operations must be defined as constants using iota
//
//--Notes:
//* Your program is complete when it compiles and prints the correct results

package main

import "fmt"

//* Operations required:
//  - Add, Subtract, Multiply, Divide
const (
	Add = iota
	Subtract
	Multiply
	Divide
)

type Operation int

//* Write a receiver function that performs the mathematical operation
//  on two operands
func (op Operation) calculate(a int, b int) int{
	switch op {
	case Add:
		return a + b
	case Subtract:
		return a - b
	case Multiply:
		return a * b
	case Divide:
		return a / b
	}
	panic("Unhandle Operation")
}

func main() {

	//* The existing function calls in main() represent the API and cannot be changed
	add := Operation(Add)
	fmt.Println(add.calculate(2, 2)) // = 4
	sub := Operation(Subtract)
	fmt.Println(sub.calculate(10, 3)) // = 7
	mul := Operation(Multiply)
	fmt.Println(mul.calculate(3, 3)) // = 9
	div := Operation(Divide)
	fmt.Println(div.calculate(100, 2)) // = 50
}
