package main

import (
	"fmt"
)

type Product interface {
	GetPrice() float64
	SetDiscount(discount float64)
	GetDescription() string
}

type Item struct {
	Name     string
	Price    float64
	Discount float64
}

func (i *Item) GetPrice() float64 {
	return i.Price * (1 - i.Discount)
}

func (i *Item) SetDiscount(discount float64) {
	if discount < 0 {
		i.Discount = 0
	} else if discount > 1 {
		i.Discount = 1
	} else {
		i.Discount = discount
	}
}

func (i *Item) GetDescription() string {
	return fmt.Sprintf("%s: цена %.2f (скидка %.0f%%)", i.Name, i.GetPrice(), i.Discount*100)
}

func CalculateTotalCost(products []Product) float64 {
	total := 0.0
	for _, product := range products {
		total += product.GetPrice()
	}
	return total
}

func NewItem(name string, price float64, discount float64) *Item {
	item := &Item{Name: name, Price: price}
	item.SetDiscount(discount)
	return item
}

func main() {
	products := []Product{
		NewItem("принтер", 5000.0, 0.1), // 10% скидка
		NewItem("ноутбук", 40000.0, 0.15), // 15% скидка
		NewItem("комьпютерная мышь", 500.0, 0.05),  // 5% скидка
	}

	for _, product := range products {
		fmt.Println(product.GetDescription())
	}

	totalCost := CalculateTotalCost(products)
	fmt.Printf("Общая стоимость товаров: %.2f\n", totalCost)
}