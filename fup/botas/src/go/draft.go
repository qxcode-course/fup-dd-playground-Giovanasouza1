package main
import "fmt"
func main() {
    var N int
    fmt.Scan(&N)

    direita := make([]int, 61)
	esquerda := make([]int, 61)

    for i := 0; i < N; i++ {
		var tamanho int
		var pe string

        fmt.Scan(&tamanho, &pe)

		if pe == "D" {
			direita[tamanho]++
		} else {
			esquerda[tamanho]++
		}
	}

	botas := 0

	for tamanho := 30; tamanho <= 60; tamanho++ {
		if direita[tamanho] < esquerda[tamanho] {
			botas += direita[tamanho]
		} else {
			botas += esquerda[tamanho]
		}
	}

	fmt.Println(botas)
}