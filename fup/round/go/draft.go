package main
import ("fmt"; "math")
func main() {
    var n1 float64
    var carac string
    fmt.Scan(&carac, &n1)
    if carac == "r" {
        fmt.Println(math.Round(n1))
    } else if carac == "f" {
        fmt.Println(math.Floor(n1))
    } else if carac == "c" {
        fmt.Println(math.Ceil(n1))
    }
}