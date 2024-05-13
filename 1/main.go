package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Carro struct {
	Nome   string `json:"nome"`
	Modelo string `json:"modelo"`
	Ano    int    `json:"-"`
}

func (c Carro) Andar() {
	fmt.Println("O carro " + c.Modelo + " está andando")
}

func (c Carro) Parar() {
	fmt.Println("O carro " + c.Modelo + " está parando")
}

func main() {
	carro1 := Carro{Nome: "Fiat", Modelo: "Uno", Ano: 2022}
	carro2 := Carro{Nome: "Ford", Modelo: "Fiesta", Ano: 2022}
	carro1.Andar()
	carro1.Parar()
	carro2.Andar()
	carro2.Parar()
	fmt.Println(carro1.Modelo)
	fmt.Println(carro2.Modelo)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(carro1)
	})
	http.ListenAndServe(":8080", nil)
}
