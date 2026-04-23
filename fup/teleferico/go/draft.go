package main
import ("fmt"; "math")
func main() {
    var c, a float64
    fmt.Scan(&c, &a)
    resto := a % c
    if a % c > 0 {
        fmt.Println(math.Ceil(resto))
    }
}
