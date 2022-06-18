package manager

import (
	"fmt"
)

type Manager struct {
	name     string
	password string
}

func OurBoss() {
	b := Manager{}
	NameBoss := "Ibrohim"
	Password := "1234321"
	fmt.Print("Enter name of Boss: ")
	fmt.Scan(&b.name)
	fmt.Print("Enter password of Boss: ")
	fmt.Scan(&b.password)
	if b.name == NameBoss && b.password == Password {
		x := 0
		fmt.Print(`
			1 - Show products
			2 - Add products
			3 - Increase
			4 - Delete
			5 - Exit		
`)
		fmt.Scan(&x)
		if x == 1 {
			fmt.Println("hello")
		} else if x == 2 {
			fmt.Println("world")
		} else if x == 5 {

		}
	}
}
