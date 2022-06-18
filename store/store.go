package store

import (
	"fmt"
	"shop/basket"
	"shop/inventory"
	"shop/product"
)

type Store struct {
<<<<<<< HEAD
=======
	Inventory inventory.Inventory
>>>>>>> main
	Basket    basket.Basket
	Profit    uint
	Budget    uint
}

func (s Store) Sell() {
	s.Budget = 20000000
	for {
<<<<<<< HEAD
		inventory.Show()
		fmt.Print("Enter product name >>> ")
		_, err := fmt.Scan(&s.Basket.Name)
		if err != nil{
			
		}
		s.Basket.Quantity, s.Basket.Price, s.Basket.OriginPrice := inventory.Check(s.Basket.Name)
=======
		product.Show()
		name := ""
		var quantity uint
		fmt.Print("Enter product name to buy --> ")
		_, err := fmt.Scan(&name)
		if err != nil {
			fmt.Println("can't can't read input")
			continue
		}
		quantity, boolean := inventory.Check(name)
		AGAIN:
		if boolean {
			fmt.Print("Enter product quantity to buy --> ")
			_, err = fmt.Scan(&quantity)
			if err != nil {
				fmt.Println("can't read input")
				continue
			}
		} else {
			boolean = inventory.Inventory.AddProduct(name, s.Budget)
			if boolean {
				goto AGAIN
			}else{
				fmt.Println("we have not enough budget\nWill you buy something else[y/n] --> ")
				check := ""
				fmt.Scan(&check)
			}
		}
>>>>>>> main
	}
}
