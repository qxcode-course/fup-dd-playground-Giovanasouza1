package main

import "fmt"

func main() {
	var n, x, y int
	fmt.Scan(&n, &x, &y)
	var c string
	fmt.Scan(&c)
	var s int
	fmt.Scan(&s)
	if c == "R" {
		fmt.Println(((x + s) % n), y)
	} else if c == "U" {
		fmt.Println(x, ((y - s + n) % n))
	} else if c == "L" {
		if ((x - s + n) % n) < 0 {
				fmt.Println((((x - s + n) % n)*-1), y)
		} else {
			fmt.Println(((x - s + n) % n), y)
		}
	} else if c == "D" {
		fmt.Println(x, ((y + s) % n))
	}
}
