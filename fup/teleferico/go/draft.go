package main
import ("fmt"; "math")
func main() {
    var c, a float64
    fmt.Scan(&c, &a)
    viagens := a / (c - 1)
    fmt.Println(math.Ceil(viagens))
}
