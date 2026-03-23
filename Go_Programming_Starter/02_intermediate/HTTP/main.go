package main

import (
	"encoding/json"
	"fmt"
)

type spouse struct {
	Name string
	Age  int
}

type person struct {
	Name     string
	Age      int
	Married  bool
	Spouse   spouse
	Children []string
}

func main() {
	jsonData := []byte(`
		{
			"Name":"Sundra Bomjan",
			"Age":90,
			"Married":false,
			"Spouse":{
				"Name":"IDK",
				"Age":10000
			},
			"Children":["Su", "Bo"]
		}
	`)

	var p person

	err := json.Unmarshal(jsonData, &p)

	if err != nil {
		fmt.Println("Error Found: ", err)
	} else {
		fmt.Println("Data decoded: ", p)
	}

}
