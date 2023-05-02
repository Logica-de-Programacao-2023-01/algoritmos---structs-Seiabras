package main

import "fmt"

//Crie uma struct chamada Retângulo com os campos "largura" e "altura". Escreva uma
//função que receba um Retângulo como parâmetro e calcule a área do retângulo (área =
//largura * altura).

type Retangulo struct {
	largura float64
	altura  float64
}

func area(r Retangulo) {
	fmt.Println(r.largura * r.altura)
}

func main() {
	r := Retangulo{largura: 10.0, altura: 10.0}
	area(r)
}
