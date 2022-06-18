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

func Check(name string) (uint, uint, uint) {
	for _, value := range ProductList {
		if value.Name == name {
			return value.Quantity, value.Price, value.OriginPrice
		}
	}
	return 0, 0, 0
}

func WriteToData() {

}


func (i Inventory) AddProduct(name string, budget, quantity uint) (bool, uint, uint, uint) {
	p := product.Product
	d := dealer.Dealer
	p.Quantity, p.OriginPrice, budget := d.AddProduct(quantity, budget)
	p.Name = name
	p.Price = p.OriginPrice * (1/5 + 1)
	p.ProductList = append(p.ProductList, p)
	return true, p.Price, p.OriginPrice, budget
func (i Inventory) AddProduct(name string, budget uint) bool {

	p := product.Product
	d := dealer.Dealer
	for _, val := i.ProductList {
		if val.Name == name {
			
		}
	}

}

func (i Inventory) IncreaseProduct(name string, quantity, budget uint) (uint, bool) {
	p := product.Product
	d := dealer.Dealer
	p.Name = name
	p.OriginPrice, p.Quantity, a = d.IncreaseProduct(quantity, budget)
	if a == true {
		p.Quantity -= quantity
		p.Price = p.OriginPrice * (1/5+1)
		i.ProductList = append(i.ProductList, p)
		return budget, true
	}
	return budget, false
}

