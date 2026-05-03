package main
import "fmt"
func main(){
    var a, b int
    fmt.Scan(&a, &b)
    acc := 0
    if a > b {
        fmt.Println("invalido")
        return
    }
    for i := a; i <= b; i++ {
        if i % 2 == 0 && i % 3 ==0 {
            acc += i
        }
    }
    fmt.Println(acc)
}
