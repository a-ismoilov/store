package store

import (
	"fmt"
	"shop/basket"
	"shop/inventory"
	"shop/product"
)

type Store struct{
	Inventory inventory.Inventory
	Basket basket.Basket
	Profit uint
	Budget uint
}

func (s Store) Sell(){
	s.Budget = 20000000
	for{
		product.Show()
		name := ""
		quantity := 0
		fmt.Print("Enter product name to buy --> ")
		_, err := fmt.Scan(&name)
		if err != nil{
			fmt.Println("can't find product")
			continue
		}
		quantity, boolean := inventory.Check(name)
		fmt.Print("Enter product quantity to buy --> ")
		_, err = fmt.Scan(&quantity)
		if err != nil{
			fmt.Println("can't ")
			continue
		}
	}
}