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
	inv := i.Inventory{}

	budget := uint(20000000)
	profit := uint(0)
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
			originPrice uint
			price       uint
			quantity    uint
			name        string
			quantityCh  uint
			basket      b.Basket
		)

		fmt.Print("Enter product name >>> ")
		_, err := fmt.Scan(&name)
		if err != nil {
			fmt.Println("can't read input")
			continue
		}

		fmt.Print("Enter product quantity >>> ")
		_, err = fmt.Scan(&quantityCh)
		if err != nil {
			fmt.Println("can't read input")
			continue
		}

		quantity, price, originPrice = i.Check(name)

		if quantity != 0 {
			if quantityCh-quantity >= 0 {
				basket = b.Basket{
					Name:     name,
					Price:    price * quantity,
					Quantity: quantity,
				}
				profit = price*quantity - originPrice*quantity
				b.Add(basket)
				continue
			} else {
				bud, boolean := inv.IncreaseProduct(name, quantity, budget)
				budget = bud
				if boolean {
					profit = price*quantity - originPrice*quantity
					basket = b.Basket{
						Name:     name,
						Price:    price * quantity,
						Quantity: quantity,
					}
					b.Add(basket)
					continue
				} else {
					fmt.Println("we can't effort with this product")
					continue
				}
			}
		} else {
			boolean, price, originPrice, bud := inv.AddProduct(name, budget, quantity)
			budget = bud
			if boolean {
				basket = b.Basket{
					Name:     name,
					Price:    price * quantity,
					Quantity: quantity,
				}
				profit = price*quantity - originPrice*quantity
				b.Add(basket)
				fmt.Print(profit)
				continue
			} else {
				fmt.Println("we can't effort with this product")
				continue
			}
		}
	}
}

func clean() {
	if err := recover(); err != nil {
		fmt.Print("you can't cheat us bro :)")
	}
}
