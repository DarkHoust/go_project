package main

type Menu struct {
	Menu []MenuItem
}

type MenuItem struct {
	Name  string
	Price int
}

var menuInstance *Menu

func getMenu() *Menu {
	if menuInstance == nil {
		menuInstance = &Menu{Menu: []MenuItem{
			{Name: "Pepperoni Pizza", Price: 1700},
			{Name: "Cheese Pizza", Price: 1500},
			{Name: "4 Season Pizza", Price: 1900},
			{Name: "Coca-cola", Price: 600},
			{Name: "Water", Price: 400},
		}}
	}
	return menuInstance
}

func (m *Menu) getPrice(name string) int {
	for _, i := range m.Menu {
		if i.Name == name {
			return i.Price
		}
	}
	return 0
}
