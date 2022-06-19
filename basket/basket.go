package basket

import (
	"fmt"
)

var baskets []Basket

type Basket struct {
	Name     string
	Price    uint
	Quantity uint
}

func Add(basket Basket) {
	baskets = append(baskets, basket)
}

func Over() {
	var sum uint
	for i := range baskets {
		sum += baskets[i].Price
		fmt.Println(baskets[i].Name, baskets[i].Price, baskets[i].Quantity)
	}
	fmt.Println("Total:", sum)
}
