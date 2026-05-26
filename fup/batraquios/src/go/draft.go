package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)
	primeiro := make([]int, N)
    for i := 0; i < N; i++ {
		fmt.Scan(&primeiro[i])
	}
    var M int
	fmt.Scan(&M)
	segundo := make([]int, M)
	for j := 0; j < M; j++ {
		fmt.Scan(&segundo[j])
	}


	lista1 := 0
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			if primeiro[i] == segundo[j] {
				lista1++
				break
			}
		}
	}

	if lista1 == N {
		fmt.Println("sim")
	} else {
		fmt.Println("nao")
	}
}
