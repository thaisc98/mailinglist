//--Summary:
//  Create a program that directs vehicles at a mechanic shop
//  to the correct vehicle lift, based on vehicle size.
//
//--Requirements:
//* The shop has lifts for multiple vehicle sizes/types:
//  - Motorcycles: small lifts
//  - Cars: standard lifts
//  - Trucks: large lifts



//
//--Notes:
//* Use any names for vehicle models

package main

import "fmt"

const (
	SmallLift = iota
	StandardLift
	LargeLift
)

type Lift int

type LiftPicker interface {
	PickLift() Lift
}

type Truck string
type Car string
type Motorcycle string

//* Vehicles have a model name in addition to the vehicle type:
//  - Example: "Truck" is the vehicle type, "Road Devourer" is a model name
func (t Truck) String() string {
	return fmt.Sprintf("Trucks: %v", string(t))
}

func (m Motorcycle) String() string {
	return fmt.Sprintf("Motorcycle: %v", string(m))
}

func (c Car) String() string {
	return fmt.Sprintf("Cars: %v", string(c))
}

func (m Motorcycle) PickLift() Lift {
	return SmallLift
}

func (c Car) PickLift() Lift {
	return StandardLift
}

func (t Truck) PickLift() Lift {
	return LargeLift
}

//* Write a single function to handle all of the vehicles
//  that the shop works on.
func sendToLift(p LiftPicker){
	switch p.PickLift(){
	case SmallLift:
		fmt.Printf("send %v to small lift \n", p)
	case StandardLift:
		fmt.Printf("send %v to standard lift \n", p)
	case LargeLift:
		fmt.Printf("send %v to large lift \n", p)
	}
}

func main() {
	
	car := Car("Sporty")
	truck := Truck("MountainCrusher")
	motorcycle := Motorcycle("CRX")

	//* Direct at least 1 of each vehicle type to the correct
	//  lift, and print out the vehicle information.
	sendToLift(car)
	sendToLift(truck)
	sendToLift(motorcycle)
}

