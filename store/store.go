package store

import (
	"fmt"
	b "shop/basket"
	i "shop/inventory"
	p "shop/product"
)

type Store struct {
	Profit uint
	Budget uint
}

func Sell() {

	var (
		s Store = Store{
			Profit: 0,
			Budget: 2000000,
		}
		inv i.Inventory
	)

	for {
		fmt.Print(`
			1 - Start/Continue
			2 - Exit
		`)
		cmd := ""
		print("choose >>> ")
		fmt.Scan(&cmd)
		switch cmd {
		case "1":
			p.Show()
		case "2":
			b.Over()
		default:
			clean()
			panic("")
		}

		var (
			originPrice uint     //original orice of the product
			price       uint     //price of the product in store
			quantity    uint     //overall quantity of theproduct
			name        string   //name of the product
			quantityIn  uint     //quantity of product that user inputs
			basket      b.Basket //basket for one product
		)

		fmt.Print("Enter product name >>> ")
		_, err := fmt.Scan(&name)
		if err != nil {
			fmt.Println("can't read input")
			continue
		}

		fmt.Print("Enter product quantity >>> ")
		_, err = fmt.Scan(&quantityIn)
		if err != nil {

			_, err = fmt.Scan(&quantityIn)
			if err != nil {

				fmt.Println("can't read input")
				continue
			}

			quantity, price, originPrice = i.Check(name)

			if quantity != 0 {
				if quantity-quantityIn >= 0 {
					basket = b.Basket{
						Name:     name,
						Price:    price * quantityIn,
						Quantity: quantityIn,
					}
					s.Profit = price*quantityIn - originPrice*quantityIn
					b.Add(basket)
					continue
				} else {
					bud, boolean := inv.IncreaseProduct(name, quantityIn, s.Budget)
					s.Budget = bud
					if boolean {
						s.Profit = price*quantityIn - originPrice*quantityIn
						basket = b.Basket{
							Name:     name,
							Price:    price * quantityIn,
							Quantity: quantityIn,
						}
						b.Add(basket)
						continue
					} else {
						fmt.Println("we can't effort with this product")
						continue
					}
				}
			} else {
				boolean, price, originPrice, bud := inv.AddProduct(name, s.Budget, quantity)
				s.Budget = bud
				if boolean {
					basket = b.Basket{
						Name:     name,
						Price:    price * quantityIn,
						Quantity: quantityIn,
					}
					s.Profit = price*quantityIn - originPrice*quantityIn
					b.Add(basket)
					fmt.Print(s.Profit)
					continue
				} else {
					fmt.Println("we can't effort with this product")
					continue
				}
			}
		}
	}
}

func clean() {
	if err := recover(); err != nil {
		fmt.Print("you can't cheat us bro :)")
	}
}
