package main
import "fmt"
func main() {
    var a, z int
    fmt.Scan(&a, &z)
    for i := a; i <= z; i ++ {
        if i % 3 == 0 && i % 5 != 0{
            fmt.Println("zig") 
        } else if i % 5 == 0 && i % 3 != 0{
            fmt.Println("zag")
        } else if i % 3 == 0 && i % 5 == 0 {
            fmt.Println("zigzag")
        } else if i % 3 != 0 {
            fmt.Println(i)
        } else if i % 5 != 0 {
            fmt.Println(i)
        } 
    }
}
