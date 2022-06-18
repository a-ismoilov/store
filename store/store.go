package store

import (
	"fmt"
	"shop/basket"
	"shop/inventory"
	"shop/product"
)

type Store struct {
	Basket    basket.Basket
	Profit    uint
	Budget    uint
}

func (s Store) Sell() {
	s.Budget = 20000000
	for {
		inventory.Show()
		fmt.Print("Enter product name >>> ")
		_, err := fmt.Scan(&s.Basket.Name)
		if err != nil{
			fmt.Println("can't read data")
		}
		s.Basket.Quantity, s.Basket.Price, s.Basket.OriginPrice := inventory.Check(s.Basket.Name)
	}
}
