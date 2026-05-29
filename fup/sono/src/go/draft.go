package main

import "fmt"

func main() {
    var h1, m1, s1 int
    var h2, m2, s2 int

    fmt.Scan(&h1, &m1, &s1, &h2, &m2, &s2)

    tempod := h1*3600 + m1*60 + s1
    tempoa := h2*3600 + m2*60 + s2

    total := tempoa - tempod

    if total < 0 {
        total += 24 * 3600
    }

    horas := total / 3600
    resto := total % 3600

    min := resto / 60
    seg := resto % 60

    fmt.Printf("%02d %02d %02d\n", horas, min, seg)
}