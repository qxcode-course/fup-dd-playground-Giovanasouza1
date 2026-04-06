package main
import "fmt"
func main() {
    var a string
    var b int
    fmt.Scan(&a, &b)
    if b < 12 {
        fmt.Print(a , " eh crianca\n")
    } else if b >= 12 && b < 18 {
        fmt.Print(a , " eh jovem\n")
    } else if b >= 18 && b < 65 {
        fmt.Print(a , " eh adulto\n")
    } else if b >= 65 && b < 1000 {
        fmt.Print(a , " eh idoso\n")
    } else {
        fmt.Print(a , " eh mumia\n")}

}
