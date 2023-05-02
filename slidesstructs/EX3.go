package main

import "fmt"

//Crie uma struct chamada Aluno com os campos "nome", "idade" e "notas". O campo
//"notas" deve ser um slice de float64. Escreva uma função que receba um Aluno como
//parâmetro e calcule a média das suas notas.

type Aluno struct {
	nome  string
	idade int
	notas []float64
}

func calcularMedia(al Aluno) float64 {
	total := 0.0
	for _, nota := range al.notas {
		total += nota
	}
	media := total / float64(len(al.notas))
	return media
}

func main() {
	aluno := Aluno{
		nome:  "Pedro",
		idade: 20,
		notas: []float64{8.0, 7.0, 4.0},
	}
	media := calcularMedia(aluno)
	fmt.Printf("A média das notas do aluno %s é: %.2f\n", aluno.nome, media)

}
