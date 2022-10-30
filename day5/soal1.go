package main

import (
	items "assignments/data_type"
	"assignments/db"
	"encoding/json"
	"fmt"
)

func main() {
	// Import Data
	var dummydb items.Items
	var rawdb = db.Dummydb
	budget := 100000
	price := 10000

	err := json.Unmarshal([]byte(rawdb), &dummydb)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Assignments
	fmt.Printf("Barang yg bisa dibeli dg Rp%.3f antara lain:", float32(budget)/1000)
	for _, v := range dummydb.PurchasableAt(budget) {
		fmt.Println(v.Name)
	}

	fmt.Printf("\nBarang termurah: %s\n", dummydb.LowestPrice())
	fmt.Printf("\nBarang termahal: %s\n", dummydb.HighestPrice())

	fmt.Printf("\nBarang dg harga Rp%.3f antara lain:", float32(price)/1000)
	for _, v := range dummydb.PriceAt(price) {
		fmt.Println(v.Name)
	}
}
