package main

import "fmt"

type Preparer interface {
	PrepareDish()
}

type Chicken string
type Salad string

func (c Chicken) PrepareDish(){
	fmt.Println("cook chicken")
}


func (c Salad) PrepareDish(){
	fmt.Println("shop salad")
	fmt.Println("add dressing")
}

func prepareDishes(dishes []Preparer){
	fmt.Println(" Prepare dishes:")
	for i := 0; i < len(dishes); i++ {
		dish := dishes[i]
		fmt.Println("---Dish: %v--\n", dish)
		dish.PrepareDish()
	}

	fmt.Println()
}

func main() {
	dishes := []Preparer{Chicken("Chicken Wings"), Salad("Tomato Salad")}
	prepareDishes(dishes)
}
