package main
import "fmt"
func main() {
    var soma int
    fmt.Scan(&soma)
    alfabeto := [26]string{
        "a","b","c","d","e","f","g","h","i","j",
    "k","l","m","n","o","p","q","r","s","t",
    "u","v","w","x","y","z",
}
    
   
for i := 0; i < 26; i ++ {
    if i + 1 == soma {
            fmt.Println(alfabeto[i])
        }
    } 
if soma == 0 {
    fmt.Println("joguem de novo")
}
if soma > 26 {
    soma = soma - 26 
    fmt.Println(alfabeto[i])

}
}