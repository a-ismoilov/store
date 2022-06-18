package inventory

import (
	"hash/adler32"
	"runtime/trace"
	"shop/dealer"
	"shop/product"
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

func (i Inventory) AddProduct(name string, budget uint) bool {
	
}

func (i Inventory) IncreaseProduct(name string, quantity uint, price uint, budget uint) uint {
	p := product.Product
	d := dealer.Dealer
	p.Name = name
	p.OriginPrice, p.Quantity, a = d.IncreaseProduct(quantity, budget)
	if a == true {
		p.Quantity -= quantity
		p.Price = p.OriginPrice * (1/5+1)
		i.ProductList = append(i.ProductList, p)
		return budget
	}
}

