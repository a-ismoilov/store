package dealer

import (
	"math/rand"
	"time"
)

type Dealer struct{}

func (d Dealer) AddProduct(quant uint, price uint, budget uint) (uint, uint, uint) {
	orignPrice := RandPrice(quant, price)
	if orignPrice*quant > budget {
		for budget != 0 {
			budget -= orignPrice
			quant += 1
		}
	}
	return quant, orignPrice, budget
}
func IncreaseProduct(quant uint, orgprice uint, budget uint) {

}
func RandPrice(quant uint, price uint) uint {
	rand.Seed(time.Now().Unix())
	var p uint
	if quant > 1 && quant < 10 {
		p = price
	} else if quant >= 10 && quant < 50 {
		r := uint(rand.Intn(10) + 1)
		p = price + r
	} else {
		r := uint(rand.Intn(30) + 1)
		p = price + r
	}
	return p
}
