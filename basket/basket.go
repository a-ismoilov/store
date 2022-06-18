package basket

import(
	"fmt"
)

type baskets []Basket

type Basket struct{
	Name string
	Price uint
	Quantity uint
}

func ListCustomProduct(){
	sum := 0
	for i := range baskets{
		sum += baskets[i].Price
		fmt.Println(baskets[i].Name,baskets[i].Price,baskets[i].Quantity)
	}
	fmt.Println("sum:",sum)
}
