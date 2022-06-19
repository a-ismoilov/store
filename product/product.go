package main

import (
	"fmt"
	"io"
	"os"
)

type Product struct {
	Name        string
	Quantity    uint
	Price       uint
	OriginPrice uint
}

func main() {
	file, err := os.OpenFile("data.txt", os.O_RDWR|os.O_CREATE, 0600)
	if err != nil {
		fmt.Print("|can't_open|", err)
	}

	data, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("|can't_read|")
	}
	fmt.Println(string(data))
	file.Close()
}

func Show(){

}
