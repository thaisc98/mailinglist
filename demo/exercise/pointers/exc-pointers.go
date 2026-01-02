//--Summary:
//  Create a program that can activate and deactivate security tags
//  on products.
//
//--Requirements:

//* Perform the following:


package main

import "fmt"

//* Create a structure to store items and their security tag state
const ( 
	Active = true
	Inactive = false
)

//  - Security tags have two states: active (true) and inactive (false)
type SecurityTag bool 

type Item struct {
	name string
	tag SecurityTag
}
//* Create functions to activate and deactivate security tags using pointers
func activate(tag *SecurityTag){
	*tag = Active 
}
func deactivate(tag *SecurityTag){
	*tag = Inactive 
}

//* Create a checkout() function which can deactivate all tags in a slice
func checkout(items []Item){
	fmt.Println("Checking out...")
	for i := 0; i< len(items); i++ {
		deactivate(&items[i].tag)
	}
}
func main() {

	//  - Create at least 4 items, all with active security tags
	shirt := Item{"shirt", Active}
	pants := Item{"pants", Active}
	purse := Item{"purse", Active}
	watch := Item{"watch", Active}

	//  - Store them in a slice or array
	items := []Item{shirt,pants,purse,watch}
	fmt.Println("Initial", items)

	//  - Deactivate any one security tag in the array/slice
	deactivate(&items[0].tag)
	
	fmt.Println("Item 0 deactivated", items)
	//  - Call the checkout() function to deactivate all tags
	checkout(items)
	//  - Print out the array/slice after each change
	fmt.Println("check out", items)
	
	activate(&items[1].tag)
	fmt.Println("Final ", items)
}

