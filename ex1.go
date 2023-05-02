package main

import (
	"fmt"
	"math"
)

// Circulo Crie uma struct chamada Circulo com o campo "raio".
// Escreva uma função que receba um Circulo como parâmetro e
// calcule a área do círculo (área = pi * raio * raio).
// Dica: use a constante math.Pi para representar o número pi.
type Circulo struct {
	raio float64
}

func area(a Circulo) {
	fmt.Println(math.Pow(a.raio, 2) * math.Pi)
}

func main() {
	ar := Circulo{raio: 12.0}
	area(ar)

}
