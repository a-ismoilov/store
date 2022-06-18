package inventory

import (
	"runtime/trace"
	"shop/dealer"
)

type Inventory struct {
	ProductList product.ProductList
}

func Check(name string) (uint, bool) {
	for i, value := range ProductList {
		if value.Name == name {
			return value.Quantity, true
		}
	}
	return 0, false
}

func WriteToData() {

}

func (i Inventory) AddProduct(name string, quantity uint, price uint, budget uint) bool {
	p := product.Product
	d := dealer.Dealer
	p.Name = name
	p.Quantity, p.OriginPrice, budget := d.AddProduct(quantity, price, budget)
	p.Price = p.OriginPrice * (1/5 + 1)
	if p.Quantity > quantity {
		p.Quantity -= quantity
		i.ProductList = append(i.ProductList, p)
		WriteToData()
		return true
	}
	return false
}

func (i Inventory) IncreaseProduct(name string, quantity uint, budget uint) uint {
	return uint(0)
}

