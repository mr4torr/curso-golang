package main

import (
	"encoding/json"
	"fmt"
)

type product struct {
	ID    int      `json:"id"`
	Name  string   `json:"name"`
	Price float64  `json:"price"`
	Tags  []string `json:"tags"`
}

func main() {

	// struct para json
	p1 := product{1, "Notebook", 5999.56, []string{"Notebook", "I5", "Intel"}}

	p1Json, _ := json.Marshal(p1)
	fmt.Println(string(p1Json))

	// json para struct
	var p2 product
	jsonString := `{"id":1,"name":"Caneta","price":9.56,"tags":["Escritorio","azul"]}`

	json.Unmarshal([]byte(jsonString), &p2)
	fmt.Println(p2)
	fmt.Println(p2.Tags[1])
}
