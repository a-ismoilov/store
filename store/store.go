package store

import (
	"fmt"
	"shop/basket"
	"shop/inventory"
	"shop/product"
)

type Store struct {
	inventory inventory.Inventory
	Profit    uint
	Budget    uint
}

func (s Store) Sell() {
	s.Budget = 20000000
	for {
		var (
			originPrice uint
			price uint
			quantity uint
			name string
			quantityCh uint
			budget uint
			basket basket.Basket
		)
		// inventory.Show()
		fmt.Print("Enter product name >>> ")
		_, err := fmt.Scan(&name)
		if err != nil{
			fmt.Println("can't read data")
		}
		_, err = fmt.Scan(&quantityCh)
		quantity,price, originPrice = inventory.Check(name)
		if quantity != 0{
			if quantityCh - quantity >= 0{
				basket.Name = name
				basket.Price = price * quantity
				basket.Quantity = quantity
				s.Profit = price*quantity-originPrice*quantity
			}else{
				budget,boolean := s.inventory.IncreaseProduct(name,quantity,budget)
				budget = budget
				if boolean{
					basket.Name = name
					basket.Price = price * quantity
					basket.Quantity = quantity
					s.Profit = price*quantity-originPrice*quantity
				}else{
					fmt.Println("we can't effort with this product")
					continue
				}
			}
		}else{
			boolean, price, originPrice, budget := s.inventory.AddProduct(name,budget,quantity)
			price, originPrice, budget = price, originPrice, budget
			if boolean{
				basket.Name = name
				basket.Price = price * quantity
				basket.Quantity = quantity
				s.Profit = price*quantity-originPrice*quantity
				basket.Add()
			}else{
				fmt.Println("we can't effort with this product")
				continue
			}
		}
	}
}
