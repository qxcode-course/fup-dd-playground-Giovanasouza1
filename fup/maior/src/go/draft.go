package main

import "fmt"

func main() {
    var chute, valor float32
    var escolha string

    fmt.Scan(&chute, &valor, &escolha)

    if chute > valor && escolha == "M" {
        fmt.Println("primeiro")
    } else if chute > valor && escolha == "m" {
        fmt.Println("segundo")
    } else if chute < valor && escolha == "M" {
        fmt.Println("segundo")
    } else if chute < valor && escolha == "m" {
        fmt.Println("primeiro")
    } else if chute == valor && escolha == "M" {
        fmt.Println("primeiro")
    } else if chute == valor && escolha == "m" {
        fmt.Println("segundo")
    }
}