package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to Dodo Pizzeria!")
	fmt.Println("What do you want to order?:")
	fmt.Println("1. Pizza")
	fmt.Println("2. Drink")
	fmt.Println("3. Other")
	fmt.Println("q. Done")

	cart := NewCart()

	for {
		var option string
		fmt.Scanln(&option)

		switch option {
		case "1":
			fmt.Println("Select a pizza from the menu:")
			menu := getMenu()
			for _, item := range menu.Items {
				fmt.Printf("%s - $%d\n", item.Name, item.Price)
			}

			var pizzaChoice string
			fmt.Scanln(&pizzaChoice)

			// Check if the selected pizza is in the menu
			pizzaPrice := menu.getPrice(pizzaChoice)
			if pizzaPrice == 0 {
				fmt.Println("Invalid pizza selection.")
				continue
			}

			// Add the selected pizza to the cart
			cart.AddItem(MenuItem{Name: pizzaChoice, Price: pizzaPrice})
			fmt.Printf("%s added to your cart.\n", pizzaChoice)

		case "2":
			// Order Drink
			// Implement similar logic as pizza for ordering drinks

		case "3":
			// Order Other items
			// Implement logic to order other items

		case "q":
			fmt.Println("Order Summary:")
			for _, item := range cart.items {
				fmt.Printf("%s - $%d\n", item.Name, item.Price)
			}
			totalCost := cart.GetCost()
			fmt.Printf("Total Cost: $%d\n", totalCost)

			// Choose delivery option
			fmt.Println("Select a delivery option:")
			fmt.Println("1. Deliver to Home")
			fmt.Println("2. Pick up from Pizzeria")
			var deliveryOption string
			fmt.Scanln(&deliveryOption)

			var deliveryMethod Delivery
			switch deliveryOption {
			case "1":
				deliveryMethod = &DeliverToHome{}
			case "2":
				deliveryMethod = &PickUpDeliver{}
			default:
				fmt.Println("Invalid delivery option.")
				continue
			}

			deliveryMethod.Deliver()

			fmt.Println("Thank you for ordering from Dodo Pizzeria!")
			return

		default:
			fmt.Println("Invalid option. Please choose a valid option.")
		}
	}
}
