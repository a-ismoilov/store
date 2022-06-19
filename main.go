package main

import (
	"fmt"
	p "shop/basket"
	m "shop/manager"
	s "shop/store"
)

func main() {
	clean()
	for {
		fmt.Print(`
		1 - shopping
		2 - manager
		3 - exit
		`)

		cmd := ""
		_, err := fmt.Scan(&cmd)
		if err != nil {
			fmt.Print("can't read input")
			continue
		}
		
		switch cmd {
		case "1":
			s.Sell()
		case "2":
			m.OurBoss()
		case "3":
			p.Over()
			return
		default:
			panic("what did you do bro?")
		}
	}
}

func clean() {
	if err := recover(); err != nil {
		p.Over()
		return
	}
}
