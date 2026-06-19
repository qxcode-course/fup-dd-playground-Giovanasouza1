package main

import "fmt"

func main() {
	var qtd int
	fmt.Scan(&qtd)

	vetor := make([]int, qtd)

	for i := 0; i < qtd; i++ {
		fmt.Scan(&vetor[i])
	}

	maxfreq, atual := 1, 1

	// Descobre a maior frequência
	for i := 1; i < qtd; i++ {
		if vetor[i] == vetor[i-1] {
			atual++

			if atual > maxfreq {
				maxfreq = atual
			}
		} else {
			atual = 1
		}
	}

	// Mostra os números com frequência máxima
	fmt.Print("[ ")
	atual = 1

	for i := 1; i <= qtd; i++ {
		if i < qtd && vetor[i] == vetor[i-1] {
			atual++
		} else {
			if atual == maxfreq {
				fmt.Printf("%d ", vetor[i-1])
			}
			atual = 1
		}
	}

	fmt.Println("]")
}