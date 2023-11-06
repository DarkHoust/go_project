package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

// Занимается управлением всех заказов в системе

type Observer interface {
	Update(message string)
}

// User теперь реализует интерфейс Observer
type User struct {
	name     string
	phoneNum string
}

func (u *User) Update(message string) {
	fmt.Printf("Уведомление для пользователя %s: %s\n", u.name, message)
}

// OrderManager добавляем список подписчиков и методы для управления ими
type OrderManager struct {
	Orders    []Order
	Observers []Observer
}

func (om *OrderManager) AddObserver(observer Observer) {
	om.Observers = append(om.Observers, observer)
}

func (om *OrderManager) NotifyObservers(message string) {
	for _, observer := range om.Observers {
		observer.Update(message)
	}
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

// Определяет интерфейс для высчета общей стоимости заказа
type PricingStrategy interface {
	CalculateTotal(order Order) float64
}

type StandardPricingStrategy struct{}

func (sps StandardPricingStrategy) CalculateTotal(order Order, promo float64) float64 {
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
	return total * (1 - promo)
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

// Strategy implementation
type DeliverStrategy interface {
	GetDeliverOption() string
	Deliver(order *Order)
}

type HomeDeliveryStrategy struct{}

func (home HomeDeliveryStrategy) GetDeliverOption() string {
	return "1. Доставка до дома."
}

func (home HomeDeliveryStrategy) Deliver(order *Order) {
	fmt.Println("Напишите пожалуйста ваш адрес: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	order.Address = input
}

type PickUpStrategy struct{}

func (p PickUpStrategy) GetDeliverOption() string {
	return "2. Самовывоз."
}

func (p PickUpStrategy) Deliver(order *Order) {
	fmt.Println("Ваш заказ будет ждать вас!")
}

func main() {
	orderManager := GetOrderManager()

	// создаем пользователей и подписываем их на уведомления
	promo := 0.5
	fmt.Println("Добро пожаловать в наш сервис доставки пиццы!")

	var username string
	fmt.Print("Ваше Имя: ")
	_, _ = fmt.Scanln(&username)
	var phnNum string
	fmt.Print("Ваш номер телефона: ")
	_, _ = fmt.Scanln(&phnNum)
	user1 := &User{username, phnNum}
	orderManager.AddObserver(user1)
	orderManager.NotifyObservers("Скидка 50% на весь чек!")
	order := Order{Address: ""}

	fmt.Println("Menu:")
	fmt.Println("1. Пицца Пепперони (2500 тенге)")
	fmt.Println("2. Пицца Маргарита (2000 Тенге)")
	fmt.Println("3. Картошка фри (600 тенге)")
	fmt.Println("4. Картошка по деревенски (800 тенге)")
	fmt.Println("5. Каппучино (790 тенге)")
	fmt.Println("6. Латте (890 тенге)")
	fmt.Println("7. Пепси (400.00 тенге)")
	fmt.Println("8. Миринда (400.00 тенге)")
	fmt.Println("9. Севенап (400.00 тенге)")
	fmt.Println("10. Тирамису (1350 тенге)")
	fmt.Println("11. Брауни (900 тенге)")
	fmt.Print("Выберите пиццу (1-2): ")
	var pizzaChoice int
	_, _ = fmt.Scanln(&pizzaChoice)
	switch pizzaChoice {
	case 1:
		order.Pizzas = append(order.Pizzas, Pizza{Name: "Пицца Пепперони", Price: 2500})
	case 2:
		order.Pizzas = append(order.Pizzas, Pizza{Name: "Пицца Маргарита", Price: 2000})
	}

	fmt.Print("Вы хотите добавить доп.сыр? (да/нет): ")
	var extraCheese string
	_, _ = fmt.Scanln(&extraCheese)
	if extraCheese == "да" {
		extraCheeseDecorator := ExtraCheeseDecorator{}
		for i := range order.Pizzas {
			order.Pizzas[i] = extraCheeseDecorator.Decorate(order.Pizzas[i])
		}
	}

	fmt.Print("Вы хотите заказать напиток? (да/нет): ")
	var orderDrink string
	_, _ = fmt.Scanln(&orderDrink)

	if orderDrink == "да" {
		fmt.Print("Выберите напиток (7-9): ")
		var drinkChoice int
		_, _ = fmt.Scanln(&drinkChoice)
		switch drinkChoice {
		case 7:
			order.Drinks = append(order.Drinks, Drink{Name: "Пепси", Price: 400})
		case 8:
			order.Drinks = append(order.Drinks, Drink{Name: "Миринда", Price: 400})
		case 9:
			order.Drinks = append(order.Drinks, Drink{Name: "Севенап", Price: 400})

		}
	}

	fmt.Print("Выберите закуску (3-4): ")
	var snackChoice int
	_, _ = fmt.Scanln(&snackChoice)
	switch snackChoice {
	case 3:
		order.Snacks = append(order.Snacks, Snack{Name: "Картошка фри", Price: 600})
	case 4:
		order.Snacks = append(order.Snacks, Snack{Name: "Картошка по деревенски", Price: 800})
	}

	fmt.Print("Вы хотите заказать кофе? (да/нет): ")
	var orderCoffee string
	_, _ = fmt.Scanln(&orderCoffee)

	if orderCoffee == "да" {
		fmt.Print("Выберите кофе (5-6): ")
		var coffeeChoice int
		_, _ = fmt.Scanln(&coffeeChoice)
		switch coffeeChoice {
		case 5:
			order.Coffees = append(order.Coffees, Coffee{Name: "Каппучино", Price: 790})
		case 6:
			order.Coffees = append(order.Coffees, Coffee{Name: "Латте", Price: 890})
		}
	}

	fmt.Print("Вы хотите заказать десерт? (да/нет): ")
	var orderDessert string
	_, _ = fmt.Scanln(&orderDessert)

	if orderDessert == "да" {
		fmt.Print("Выберите десерт (10-11): ")
		var dessertChoice int
		_, _ = fmt.Scanln(&dessertChoice)
		switch dessertChoice {
		case 10:
			order.Desserts = append(order.Desserts, Dessert{Name: "Тирамису", Price: 1350})
		case 11:
			order.Desserts = append(order.Desserts, Dessert{Name: "Брауни", Price: 900})
		}
	}

	// Strategy
	fmt.Println("Какой способ получение заказа вам удобен?:")
	deliverOptions := []DeliverStrategy{HomeDeliveryStrategy{}, PickUpStrategy{}}

	for i, j := range deliverOptions {
		fmt.Println(j.GetDeliverOption())
		i += 1
	}

	var userDeliverOption int
	_, _ = fmt.Scanln(&userDeliverOption)

	if userDeliverOption >= 1 && userDeliverOption <= len(deliverOptions) {
		deliverOptions[userDeliverOption-1].Deliver(&order)
	} else {
		fmt.Println("Ошибка ввода, повторите попытку позже.")
		return
	}

	pricingStrategy := StandardPricingStrategy{}
	order.Total = pricingStrategy.CalculateTotal(order, promo)

	orderManager.AddOrder(order)

	fmt.Println("\nДетали заказа:")
	fmt.Println("Имя: ", username)
	fmt.Println("Номер телефона: ", phnNum)

	if userDeliverOption == 1 {
		fmt.Println("Способ получение заказа - Доставка: ")
		fmt.Printf("Адрес заказа: %s\n", order.Address)
	} else {
		fmt.Println("Способ получение заказа - Самовывоз. ")
	}

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
