vpackage main
import "fmt"
func main() {
    var N, M, L int
    fmt.Scan(&N, &M, &L)

    tamanho := make([]int, M)
    pe := make([]string, L)

    for i := 0; i < M; i++ {
        fmt.Scan(&tamanho[i])
    }
    for p := 0; p < L; p++ {
        fmt.Scan(&pe[p])
    }
    
    botas := 0
    for i := 0; i < M; i++ {
        if tamanho[i] == tamanho[i]  {
            botas++
}
    }
    for p := 0; p < L; p++ {
        if pe[p] == pe[p]  {
            botas++
}
}
    }