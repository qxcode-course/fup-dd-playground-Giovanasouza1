package main
import "fmt"
func main() {
    var a, b string   
    fmt.Scan(&a, &b)
    if a == "S" && b == "P" {
        fmt.Println("jog1")
    } else if a == "S" && b == "R" {
        fmt.Println("jog2")
    } else if (a == "S" && b == "S") {
        fmt.Println("empate")
    } else if a == "R" && b == "P" {
        fmt.Println("jog2")
    } else if a == "R" && b == "R" {
        fmt.Println("empate")
    } else if a == "R" && b == "S" {
        fmt.Println("jog1")
    } else if a == "P" && b == "P" {
        fmt.Println("empate")
    } else if a == "P" && b == "R" {
        fmt.Println("jog1")
    } else if a == "P" && b == "S" {
        fmt.Println("jog2")
    }
} 
