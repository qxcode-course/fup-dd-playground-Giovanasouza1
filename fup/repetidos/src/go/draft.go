package main
import "fmt"
func main() {
    var p, n int
	fmt.Scan(&p, &n)
	vetor := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&vetor[i])
	}

	quantidade := 0

	for i := 0; i < n; i++ {
		if vetor[i] == p {
			quantidade++
		}
}
fmt.Println(quantidade)
}
