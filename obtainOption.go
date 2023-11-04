package main

import "fmt"

type Delivery interface {
	Deliver()
}

type DeliverToHome struct{}

func (deliverToHome *DeliverToHome) Deliver() {
	fmt.Println("Delivering to Home")
}

type PickUpDeliver struct{}

func (pickUp *PickUpDeliver) Deliver() {
	fmt.Println("Pick up from Pizzeria")
}
