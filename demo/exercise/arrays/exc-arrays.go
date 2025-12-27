//--Summary:
//  Create a program that can store a shopping list and print
//  out information about the list.
//
//--Requirements:
//* Using an array, create a shopping list with enough room
//  for 4 products
//  - Products must include the price and the name
//* Insert 3 products into the array
//* Print to the terminal:
//  - The last item on the list
//  - The total number of items
//  - The total cost of the items
//* Add a fourth product to the list and print out the
//  information again

package main

import "fmt"


type Product struct {
	name string
	price int
}

func infoItems(products [4]Product){
	totalCost := 0
	totalProcessProducts := 0;
	for i := 0; i < len(products) ; i++{
		product := products[i]

		if(product.price == 0){
			break;
		}
		totalProcessProducts += 1
		totalCost += product.price
	}
	fmt.Println("The last item on the list: ", products[totalProcessProducts -1])
	fmt.Println("The total number of items: ", totalProcessProducts)
	fmt.Println("The total cost of items: ", totalCost)
}

func main() {
	products := [4]Product{
		{name: "Milk", price: 10},
		{name: "Water", price: 5},
		{name: "Chicken", price: 20},
	}

	fmt.Println("Shoping list processing..")
	infoItems(products)
	fmt.Println("Adding last product....")
	products[3].name = "tomato"
	products[3].price = 1

	infoItems(products)

}

