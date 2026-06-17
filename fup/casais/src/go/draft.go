package main
import "fmt"
func main() {
    var N int
    fmt.Scan(&N)

    especies := make([]int, N)
    casal := 0
    for i := 0; i < N; i++ {
        fmt.Scan(&especies[i])
    }
    
    for i := 0; i < N; i++ {
        if especies[i] < 0 {
            for j := 0; j < N; j++ {
                if especies[j] == - especies[i]{
                casal++
                especies[j] = 0
                especies[i] = 0
                break
        }
    }
    }
}
 fmt.Println(casal)
}