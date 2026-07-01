package main

import "fmt"

type Jogada struct {
	p1 int
	p2 int
}

func procurar_melhor_jogada(jogadas []Jogada) int {
	melhor := -1
	menorDif := 0

	for i := 0; i < len(jogadas); i++ {


		if jogadas[i].p1 < 10 || jogadas[i].p2 < 10 {
			continue
		}

		dif := jogadas[i].p1 - jogadas[i].p2

		if dif < 0 {
			dif = -dif
		}

		if melhor == -1 || dif < menorDif {
			melhor = i
			menorDif = dif
		}
	}
	return melhor
}

func main() {
	var n int
	fmt.Scan(&n)

	jogadas := make([]Jogada, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&jogadas[i].p1, &jogadas[i].p2)
	}

	ganhador := procurar_melhor_jogada(jogadas)

	if ganhador == -1 {
		fmt.Println("sem ganhador")
	} else {
		fmt.Println(ganhador)
	}
}