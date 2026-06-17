package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	leitor := bufio.NewReader(os.Stdin)

	frase, _ := leitor.ReadString('\n')

	vogais := ""
	consoantes := ""

	for i := 0; i < len(frase); i++ {

		if frase[i] == 'a' ||
			frase[i] == 'e' ||
			frase[i] == 'i' ||
			frase[i] == 'o' ||
			frase[i] == 'u' {

			vogais += string(frase[i])

		} else if frase[i] != ' ' && frase[i] != '\n' {

			consoantes += string(frase[i])
		}
	}

	fmt.Println(vogais)
	fmt.Println(consoantes)
}