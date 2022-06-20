package main

import (
	"fmt"
	"shop/basket"
	"shop/manager"
	"shop/store"
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
			store.Sell()
		case "2":
			manager.OurBoss()
		case "3":
			basket.Over()
			return
		default:
			panic("what did you do bro?")
		}
	}
}

func clean() {
	if err := recover(); err != nil {
		basket.Over()
		return
	}
}
