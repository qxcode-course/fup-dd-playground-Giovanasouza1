
package main
import "fmt"
    func main() {
    var n int
    fmt.Scan(&n)
    var arr []int = make([]int, n)
    for i := range arr {
        fmt.Scan(&arr[i])
    }
    for _, valor := range arr {
    fmt.Println(valor) }

    }