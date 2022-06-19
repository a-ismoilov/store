package product

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"text/tabwriter"
)

var ProductList []Product

type Product struct {
	Name        string
	Quantity    uint
	Price       uint
	OriginPrice uint
}

func Show() {
	file, err := os.OpenFile("data.txt", os.O_RDWR|os.O_CREATE, 0600)
	if err != nil {
		fmt.Print("|can't_open|", err)
	}

	defer file.Close()

	arrData := make([][]string, 0)

	file.Seek(0, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data := scanner.Text()
		var newSlice = make([]string, 0)
		newSlice = strings.Fields(data)
		arrData = append(arrData, newSlice)
	}

	var p Product
	for i := range arrData {
		p.Name = arrData[i][0]

		v, err := strconv.Atoi(arrData[i][1])
		if err != nil {
			fmt.Println("can't read file")
		}
		p.Quantity = uint(v)

		v, err = strconv.Atoi(arrData[i][2])
		if err != nil {
			fmt.Println("can't read file")
		}
		p.Price = uint(v)

		v, err = strconv.Atoi(arrData[i][3])
		if err != nil {
			fmt.Println("can't read file")
		}
		p.OriginPrice = uint(v)

		ProductList = append(ProductList, p)

	}

	w := tabwriter.NewWriter(os.Stdout, 7, 1, 1, ' ', tabwriter.AlignRight|tabwriter.Debug)
	fmt.Println("Inventory: ")
	fmt.Fprintln(w, "name\tquantity\tprice\toriginal price\t")

	for _, p := range ProductList {
		fmt.Fprintf(w, "%s\t%d\t%d\t%d\t\n", p.Name, p.Quantity, p.Price, p.OriginPrice)
	}

	w.Flush()
}
