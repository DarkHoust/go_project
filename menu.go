package main

type Menu struct {
	Items []MenuItem
}

type MenuItem struct {
	Name  string
	Price int
}

var menuInstance *Menu

func getMenu() *Menu {
	if menuInstance == nil {
		menuInstance = &Menu{Items: []MenuItem{
			{Name: "Pepperoni Pizza", Price: 1700},
			{Name: "Cheese Pizza", Price: 1500},
			{Name: "4 Season Pizza", Price: 1900},
			{Name: "Coca-cola", Price: 600},
			{Name: "Water", Price: 400},
		}}
	}
	return menuInstance
}
