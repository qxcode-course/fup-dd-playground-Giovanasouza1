
package main

import "fmt"

func media(vetor []float64) float64 {
	soma := 0.0

	for i := 0; i < len(vetor); i++ {
		soma += vetor[i]
	}

	return soma / float64(len(vetor))
}

func soldado(vetor []float64) {
	m := media(vetor)

	for i := 0; i < len(vetor); i++ {
		if vetor[i] < m {
			fmt.Print("P")
		} else if vetor[i] == m {
			fmt.Print("M")
		} else {
			fmt.Print("G")
		}

		if i < len(vetor)-1 {
			fmt.Print(" ")
		}
	}

	fmt.Println()
}

func main() {
	var qtd int
	fmt.Scan(&qtd)

	vetor := make([]float64, qtd)

	for i := 0; i < qtd; i++ {
		fmt.Scan(&vetor[i])
	}

	fmt.Printf("%.2f\n", media(vetor))
	soldado(vetor)
}