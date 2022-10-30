package main

import "fmt"

type item []struct {
	name  string
	price int
}

var structItems = item{
	{name: "Benih Lele", price: 50000},
	{name: "Pakan lele cap menara", price: 25000},
	{name: "Probiotik A", price: 75000},
	{name: "Probiotik Nila B", price: 10000},
	{name: "Pakan Nila", price: 20000},
	{name: "Benih Nila", price: 20000},
	{name: "Cupang", price: 5000},
	{name: "Benih Nila", price: 30000},
	{name: "Benih Cupang", price: 10000},
	{name: "Probiotik B", price: 10000},
}

func (items item) purchasableAt(budget int) item {
	// Sort data from lowest to highest price and
	// return highest amount of items within specified budget

	// init local var
	var result item
	var totalPrice int = 0

	// Sort array from lowest to highest price
	for i := 0; i < len(items); i++ {
		for j := (i + 1); j < len(items); j++ {
			if items[j].price < items[i].price {
				items[i], items[j] = items[j], items[i]
			}
		}
	}

	// store highest amount of items within specified budget
	for _, val := range items {
		totalPrice += val.price
		result = append(result, val)

		if totalPrice >= budget {
			break
		}
	}
	return result
}

func (items item) lowestPrice() string {
	// Find the lowest price item

	// init local var
	var itemName string = items[0].name
	var itemPrice int = items[0].price

	// Search the lowest price item
	for _, v := range items {
		if v.price < itemPrice {
			itemPrice = v.price
			itemName = v.name
		}
	}
	return itemName
}

func (items item) highestPrice() string {
	// Find the highest price item

	// init var
	var itemName string = items[0].name
	var itemPrice int = items[0].price

	// Search the highest price item
	for _, v := range items {
		if v.price > itemPrice {
			itemPrice = v.price
			itemName = v.name
		}
	}
	return itemName
}

func (items item) priceAt10k() {
	// Print items whose price equals to 10k
	for _, val := range items {
		if val.price == 10000 {
			fmt.Println(val.name)
		}
	}
}

func main() {
	structItems.priceAt10k()
	fmt.Println(structItems.highestPrice())
	fmt.Println(structItems.lowestPrice())
	fmt.Println(structItems.purchasableAt(100000))
}
