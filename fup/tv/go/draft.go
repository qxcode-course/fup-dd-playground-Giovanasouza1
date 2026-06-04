package main
import "fmt"
func main() {
	var valor, juros float64
	var parcelas int
	fmt.Scan(&valor, &parcelas)

	
	if parcelas == 2 {
		juros = 1.05
	} else if parcelas == 1 {
		juros = 1
	} else if parcelas == 3 {
		juros = 1.1
	} else if parcelas == 4 {
		juros = 1.15 
	} else if parcelas == 5 {
		juros = 1.2
	} else if parcelas == 6 {
		juros = 1.25
	}else if parcelas == 7 {
		juros = 1.30
	} else if parcelas == 8 {
		juros = 1.35
	} else if parcelas == 9 {
		juros = 1.40
	} else if parcelas == 10 {
		juros = 1.45
	}

	OPERAÇAO2 := valor * juros
	OPERAÇAO1 := OPERAÇAO2 / float64(parcelas)
	fmt.Printf("%.2f\n", OPERAÇAO1)
	fmt.Printf("%.2f\n", OPERAÇAO2)
	}
	