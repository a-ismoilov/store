package store

import (
	"shop/basket"
	"shop/inventory"
)

type Store struct{
	Inventory inventory.Inventory
	Basket basket.Basket
	Profit uint
	Budget uint
}