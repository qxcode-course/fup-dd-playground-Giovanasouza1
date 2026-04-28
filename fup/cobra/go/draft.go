package main

import "fmt"

func main() {
	var n, x, y, s int
	var c string
	fmt.Scan(&n, &x, &y, &c, &s)t
	if c == "R" || c == "L" || c == "U" || c == "D" {
		if c == "R" {
			if x+s > n {
				fmt.Println(((x + s) - n), y)
			} else {
				fmt.Println((x + s), y)
			}
		} else if c == "L" {
			if x-s < 0 {
				fmt.Println(((x + n) - s), y)
			} else {
				fmt.Println((x - s), y)
			}
		} else if c == "U" {
			if y-s < 0 {
				fmt.Println(x, ((y + n) - s))
			} else {
				fmt.Println(x, (y - s))
			}
		} else if c == "D" {
			if y+s > n {
				fmt.Println(x, ((y + s) - n))
			} else {
				fmt.Println(x, (y + s))
			}
		}
	}
}
