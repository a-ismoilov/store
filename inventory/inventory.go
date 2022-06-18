package inventory

import "shop/dealer"

type Inventory struct {}

func (i Inventory) AddProduct(name string, quantity uint)  {
	p := product.Product
	pl := product.ProductList
	d := dealer.Dealer
	p.Name = name
	p.Quantity = quantity
	p.OriginPrice = d.AddProduct(name, quantity)
	p.Price = p.OriginPrice * (1/5+1)
	pl = append(pl, p)
}