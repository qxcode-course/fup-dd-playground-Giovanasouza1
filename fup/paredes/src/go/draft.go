package main
import "fmt"
func main() {
    var N int
    fmt.Scan(&N)

    paredes := make([]int, N)

    for i := 0; i < N; i++ {
    fmt.Scan(&paredes[i])
}

cont := 0
ve := 0
    for i := 0; i < N; i++ {
        if ve < paredes[i] {
            ve = paredes[i]
            cont++
        }
    

}
fmt.Println(cont)
}