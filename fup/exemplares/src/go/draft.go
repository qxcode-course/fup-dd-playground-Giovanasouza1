package main
import "fmt"
func main() {
    var N int
    fmt.Scan(&N)

    vetor := make([]int, N)

    for i := 0; i < N; i ++ {
        fmt.Scan(&vetor[i])
    }

    unicos := []int {}

     for i := 0; i < N; i ++ {
        repetido := false
        for j := 0; j < len(unicos); j++ {
        if vetor[i] == unicos[j]  {
            repetido = true
            break
        }
        }
        if !repetido {
        unicos = append(unicos, vetor[i])
}
}
for i := 0; i < len(unicos)-1; i++ {
    for j := 0; j < len(unicos)-1 -i; j++{
    if unicos[j] > unicos[j+1] {
        unicos[j], unicos[j+1] = unicos[j+1], unicos[j]
        }
}
}
for i := 0; i < len(unicos); i++ {
		if i > 0 {
            fmt.Print(" ")
}
fmt.Print(unicos[i])
}
fmt.Println()
}
