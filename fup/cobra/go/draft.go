package main

import "fmt"

func main() {
	var n, x, y, s int
	var c string
	fmt.Scan(&n, &x, &y, &s, &c)
	if c == "R"  {
		fmt.Println(((x + s) % n), y)
	} else if c == "U" {
		fmt.Println(x, ((y - s + n) % n))
	} else if c == "L" {
		fmt.Println(((x + s) % n), y)
	} else if c == "D" {
		fmt.Println(x, ((y + s) % n))
	}
}

