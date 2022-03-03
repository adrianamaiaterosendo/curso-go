package main

import (
	"encoding/json"
	"fmt"
)

type produto struct {
	ID    int      `json:"id"`
	Nome  string   `json:"nome"`
	Preco float64  `json:"preco"`
	Tags  []string `json:"tags"`
}

func main() {
	p1 := produto{1, "Celular", 999.99, []string{"Promoção", "Eletronicos"}}
	p1Json, _ := json.Marshal(p1)

	fmt.Println(string(p1Json))

	jsonString := `{"id":2,"nome":"Maquina de Lavar","preco":1999.99,"tags":["Promoção","Eletrodomésticos"]}`
	var p2 = produto{}
	json.Unmarshal([]byte(jsonString), &p2)

	fmt.Println(p2)

}
