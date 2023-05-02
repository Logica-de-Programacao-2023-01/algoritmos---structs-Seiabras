package main

import "fmt"

// Livro Crie uma struct chamada Livro com os campos "título", "autor" e "ano de publicação".
// Escreva uma função que receba um Livro como parâmetro e imprima suas informações.

type Livro struct {
	Titulo          string
	Autor           string
	AnoDePublicacao int
}

func biblioteca(l Livro) {
	fmt.Println("Título: ", l.Titulo)
	fmt.Println("Autor: ", l.Autor)
	fmt.Println("Ano de publicação: ", l.AnoDePublicacao)
}
func main() {
	l := Livro{Titulo: "Hoje", Autor: "Carro", AnoDePublicacao: 1220}
	biblioteca(l)
}
