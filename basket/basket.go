package basket

import (
	"fmt"
)

type Basket struct {
	Name     string
	Price    uint
	Quantity uint
}

var baskets []Basket

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
