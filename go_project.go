package main

import (
	"fmt"
	"sync"
)

// Занимается управлением всех заказов в системе
type OrderManager struct {
	Orders []Order
}

type Order struct {
	Pizzas   []Pizza
	Drinks   []Drink
	Snacks   []Snack
	Coffees  []Coffee
	Desserts []Dessert
	Total    float64
	Address  string
}

// Паттерн Singleton: Создает только один экзмепляр OrderManager
var instance *OrderManager
var once sync.Once // Используем, чтобы убедиться в том, что процесс заказа был вызван только один раз

func GetOrderManager() *OrderManager {
	once.Do(func() {
		instance = &OrderManager{}
	})
	return instance
}

// Короче просто добавляем заказ в OrderManager
func (om *OrderManager) AddOrder(order Order) {
	om.Orders = append(om.Orders, order)
}

// Возвращение списка всех заказов в Менеджере
func (om *OrderManager) GetOrders() []Order {
	return om.Orders
}

// Паттерн стратегии: Определяет интерфейс для высчета общей стоимости заказа
type PricingStrategy interface {
	CalculateTotal(order Order) float64
}

type StandardPricingStrategy struct{}

func (sps StandardPricingStrategy) CalculateTotal(order Order) float64 {
	total := 0.0
	for _, pizza := range order.Pizzas {
		total += pizza.Price
	}
	for _, drink := range order.Drinks {
		total += drink.Price
	}
	for _, snack := range order.Snacks {
		total += snack.Price
	}
	for _, coffee := range order.Coffees {
		total += coffee.Price
	}
	for _, dessert := range order.Desserts {
		total += dessert.Price
	}
	return total
}

// Дкоратор: Используем, чтобы дополнять заказ пиццы дополнительными ингрендиентами
type PizzaDecorator interface {
	Decorate(pizza Pizza) Pizza
}

type ExtraCheeseDecorator struct{}

func (ecd ExtraCheeseDecorator) Decorate(pizza Pizza) Pizza {
	return Pizza{Name: pizza.Name + " с дополнительным сыром", Price: pizza.Price + 200.0}
}

type Pizza struct {
	Name  string
	Price float64
}

type Drink struct {
	Name  string
	Price float64
}

type Snack struct {
	Name  string
	Price float64
}

type Coffee struct {
	Name  string
	Price float64
}

type Dessert struct {
	Name  string
	Price float64
}

func main() {
	orderManager := GetOrderManager()

	fmt.Println("Добро пожаловать в наш сервис доставки пиццы!")

	order := Order{Address: ""}

	fmt.Println("Menu:")
	fmt.Println("1. Пицца Пепперони (2500 тенге)")
	fmt.Println("2. Пицца Пепперони (2000 Тенге)")
	fmt.Println("3. Пепси (400.00 тенге)")
	fmt.Println("4. Картошка фри (600 тенге)")
	fmt.Println("5. Картошка по деревенски (800 тенге)")
	fmt.Println("6. Каппучино (790 тенге)")
	fmt.Println("7. Латте (890 тенге)")
	fmt.Println("8. Тирамису (1350 тенге)")
	fmt.Print("Выберите пиццу (1-2): ")
	var pizzaChoice int
	fmt.Scanln(&pizzaChoice)

	fmt.Print("Вы хотите заказать напиток? (да/нет): ")
	var orderDrink string
	fmt.Scanln(&orderDrink)

	if orderDrink == "да" {
		fmt.Print("Выберите напиток (3): ")
		var drinkChoice int
		fmt.Scanln(&drinkChoice)
		switch drinkChoice {
		case 3:
			order.Drinks = append(order.Drinks, Drink{Name: "Пепси", Price: 400})
		}
	}

	fmt.Print("выберите закуску (4-5): ")
	var snackChoice int
	fmt.Scanln(&snackChoice)
	switch snackChoice {
	case 4:
		order.Snacks = append(order.Snacks, Snack{Name: "Картошка фри", Price: 600})
	case 5:
		order.Snacks = append(order.Snacks, Snack{Name: "Картошка по деревенски", Price: 800})
	}

	fmt.Print("Вы хотите заказать кофе? (да/нет): ")
	var orderCoffee string
	fmt.Scanln(&orderCoffee)

	if orderCoffee == "да" {
		fmt.Print("Выберите кофе: ")
		var coffeeChoice int
		fmt.Scanln(&coffeeChoice)
		switch coffeeChoice {
		case 6:
			order.Coffees = append(order.Coffees, Coffee{Name: "Каппучино", Price: 790})
		case 7:
			order.Coffees = append(order.Coffees, Coffee{Name: "Латте", Price: 890})
		}
	}

	fmt.Print("Вы хотите заказать десерт? (да/нет): ")
	var orderDessert string
	fmt.Scanln(&orderDessert)

	if orderDessert == "да" {
		fmt.Print("Выберите десерт (8): ")
		var dessertChoice int
		fmt.Scanln(&dessertChoice)
		switch dessertChoice {
		case 8:
			order.Desserts = append(order.Desserts, Dessert{Name: "Тирамису", Price: 1350})
		}
	}

	fmt.Print("Напишите пожалуйста ваш адрес: ")
	fmt.Scanln(&order.Address)

	switch pizzaChoice {
	case 1:
		order.Pizzas = append(order.Pizzas, Pizza{Name: "Пицца Пепперони", Price: 2500})
	case 2:
		order.Pizzas = append(order.Pizzas, Pizza{Name: "Пицца Маргарита", Price: 2000})
	}

	extraCheeseDecorator := ExtraCheeseDecorator{}
	for i := range order.Pizzas {
		order.Pizzas[i] = extraCheeseDecorator.Decorate(order.Pizzas[i])
	}
	pricingStrategy := StandardPricingStrategy{}
	order.Total = pricingStrategy.CalculateTotal(order)

	orderManager.AddOrder(order)

	fmt.Println("\nДетали заказа:")
	fmt.Printf("Адрес заказа: %s\n", order.Address)
	fmt.Println("Продукты:")
	for _, pizza := range order.Pizzas {
		fmt.Printf("- %s: Тенге %.2f\n", pizza.Name, pizza.Price)
	}
	for _, drink := range order.Drinks {
		fmt.Printf("- %s: Тенге %.2f\n", drink.Name, drink.Price)
	}
	for _, snack := range order.Snacks {
		fmt.Printf("- %s: Тенге %.2f\n", snack.Name, snack.Price)
	}
	for _, coffee := range order.Coffees {
		fmt.Printf("- %s: Тенге %.2f\n", coffee.Name, coffee.Price)
	}
	for _, dessert := range order.Desserts {
		fmt.Printf("- %s: Тенге %.2f\n", dessert.Name, dessert.Price)
	}
	fmt.Printf("В итоге: Тенге %.2f\n", order.Total)
}
