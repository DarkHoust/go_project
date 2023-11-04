package main

type Cart struct {
	items []MenuItem
}

func NewCart() *Cart {
	return &Cart{}
}

func (c *Cart) AddItem(item MenuItem) {
	c.items = append(c.items, item)
}

func (c *Cart) GetCost() int {
	totalCost := 0
	for _, i := range c.items {
		totalCost += i.Price
	}
	return totalCost
}
