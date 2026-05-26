package main
import "fmt"
func main() {
    var talbum, tbaruel int
    fmt.Scan(&talbum, &tbaruel)

    sequencia := make([]int, tbaruel)
    album := make([]int, talbum)

    for s := 0; s < tbaruel ; s++ {
        fmt.Scan(&sequencia[s])
    }

    for a := 0; a < tbaruel ; a++ {
        fmt.Scan(&album[a])
    }

    
    for s := 0; s < tbaruel - 1 ; s++ {
            if sequencia[s] == sequencia[s + 1]{
                fmt.Println(repetida)
        }
    
    }
    if sequencia[s] != sequencia[s + 1] {
        fmt.Println("N")
    }


    for i := 0; i < talbum ; i++ {
        for a := i + 1; a < talbum; a++ {
            if album[i] == album[a]{
                 fmt.Println(repete)
    
             }
        }
        }
        if repete == 0 {
            fmt.Println("N")

        }
        
       
    }
