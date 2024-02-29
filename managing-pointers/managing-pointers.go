package main

import (
	"fmt"
)

type pessoa struct {
	nome      string
	sobrenome string
	idade     int
}

func mudeMe(pes *pessoa) {
	pes.nome = "lucas"
	(*pes).idade = 30
}

func main() {
	pessoa1 := pessoa{
		nome:      "gisele",
		sobrenome: "portella",
		idade:     27,
	}
	fmt.Println(pessoa1)
	mudeMe(&pessoa1)
	fmt.Println(pessoa1)
}

