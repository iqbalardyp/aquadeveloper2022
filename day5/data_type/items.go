package items

type Items []struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
}

func (goods Items) PurchasableAt(budget int) Items {
	// Sort data from lowest to highest price and
	// return highest amount of items within specified budget

	// init local var
	var results Items
	var totalPrice int = 0

	// Sort array from lowest to highest price
	for i := 0; i < len(goods); i++ {
		for j := (i + 1); j < len(goods); j++ {
			if goods[j].Price < goods[i].Price {
				goods[i], goods[j] = goods[j], goods[i]
			}
		}
	}

	// store highest amount of items within specified budget
	for _, val := range goods {
		totalPrice += val.Price
		results = append(results, val)

		if totalPrice >= budget {
			break
		}
	}
	return results
}

func (goods Items) LowestPrice() string {
	// Find the lowest price item

	// init local var
	var itemName string = goods[0].Name
	var itemPrice int = goods[0].Price

	// Search the lowest price item
	for _, v := range goods {
		if v.Price < itemPrice {
			itemPrice = v.Price
			itemName = v.Name
		}
	}
	return itemName
}

func (goods Items) HighestPrice() string {
	// Find the highest price item

	// init var
	var itemName string = goods[0].Name
	var itemPrice int = goods[0].Price

	// Search the highest price item
	for _, v := range goods {
		if v.Price > itemPrice {
			itemPrice = v.Price
			itemName = v.Name
		}
	}
	return itemName
}

func (goods Items) PriceAt(price int) Items {
	// Print items whose price equals to 10k

	// init local var
	var results Items
	for _, val := range goods {
		if val.Price == price {
			results = append(results, val)
		}
	}
	return results
}
