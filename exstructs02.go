package main

import "fmt"

type endereco struct {
	rua    string
	numero int
	cidade string
	estado string
}

type Pessoa struct {
	nome  string
	idade int
	endereco
}

func main() {
	r := Pessoa{
		nome:  "João",
		idade: 18,
		endereco: endereco{
			rua:    "Dois",
			numero: 27,
			cidade: "Brasília",
			estado: "DF"},
	}
	fmt.Println("Nome: ", r.nome)
	fmt.Println("Idade: ", r.idade)
	fmt.Println("Estado: ", r.endereco.estado)
	fmt.Println("Cidade: ", r.endereco.cidade)
	fmt.Println("Rua: ", r.endereco.rua)
	fmt.Println("Número: ", r.endereco.numero)
}
