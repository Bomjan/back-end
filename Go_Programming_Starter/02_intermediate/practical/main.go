package main

import (
	"errors"
	"fmt"
)

type Product struct {
	Name  string
	Price float64
}

func MostExpensive(prod []Product) Product {
	me := prod[0]

	for i := 0; i < len(prod); i++ {
		if me.Price < prod[i].Price {
			me = prod[i]
		}
	}
	return me
}

func main() {
	products := []Product{
		{Name: "Apple", Price: 300.00},
		{Name: "Bread", Price: 100.00},
		{Name: "Eggs", Price: 500.00},
	}
	mostExp := MostExpensive(products)
	fmt.Println(mostExp.Name)

	data, err := fetchdata(10)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(data)
}

func fetchdata(i int) (data []int, err error) {
	if i > 9 {
		err = errors.New("This is not possivle")
		return data, err
	}
	return []int{2, 5}, nil
}
