package main

import "fmt"

func main() {
    var N, M int

    fmt.Scan(&N)
    fmt.Scan(&M)

    fig := make([]int, M)

    for i := 0; i < M; i++ {
        fmt.Scan(&fig[i])
    }

    temRepetida := false

    for i := 1; i < M; i++ {
        if fig[i] == fig[i-1] {
            if temRepetida {
                fmt.Print(" ")
            }
            fmt.Print(fig[i])
            temRepetida = true
        }
    }

    if !temRepetida {
        fmt.Print("N")
    }

    fmt.Println()

    temFaltando := false

    for num := 1; num <= N; num++ {

        encontrou := false

        for i := 0; i < M; i++ {
            if fig[i] == num {
                encontrou = true
                break
            }
        }

        if !encontrou {
            if temFaltando {
                fmt.Print(" ")
            }
            fmt.Print(num)
            temFaltando = true
        }
    }

    if !temFaltando {
        fmt.Print("N")
    }

    fmt.Println()
}