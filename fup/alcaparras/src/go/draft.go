package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	frase := scanner.Text()

	scanner.Scan()
	letra := scanner.Text()

	cont := 0

	for i := 0; i < len(frase); i++ {
		if frase[i] == letra[0] {
			cont++
		}
	}

	fmt.Println(cont)
}