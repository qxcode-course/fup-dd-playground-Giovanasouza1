package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	valor := make([]float64, n)
	chute := make([]float64, n)
	escolha := make([]string, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&valor[i])
	}

	for i := 0; i < n; i++ {
		fmt.Scan(&chute[i])
	}

	for i := 0; i < n; i++ {
		fmt.Scan(&escolha[i])
	}

	primeiro := 0
	segundo := 0

	for i := 0; i < n; i++ {

		if chute[i] == valor[i] {
			primeiro++
		} else if escolha[i] == "M" && valor[i] < chute[i] {
			primeiro++
		} else if escolha[i] == "m" && valor[i] > chute[i] {
			primeiro++
		} else {
			segundo++
		}
	}

	if primeiro > segundo {
		fmt.Println("primeiro")
	} else if segundo > primeiro {
		fmt.Println("segundo")
	} else {
		fmt.Println("empate")
	}
}