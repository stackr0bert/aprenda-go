// Escreva expressões utilizando os seguintes operadores, e atribua seus valores a variáveis:
// ==, !=, <=, <, >=, >
package main

import "fmt"

var (
	num1, num2 float64
	operacao   int
)

func main() {
	fmt.Print("Escolha um valor: ")
	_, err := fmt.Scan(&num1)
	if err != nil {
		fmt.Print("Erro ao receber o valor")
		return
	}
	fmt.Print("Escolha outro valor: ")
	_, err = fmt.Scan(&num2)
	if err != nil {
		fmt.Print("Erro ao receber o valor")
		return
	}

	fmt.Println("Agora escolha uma operação: ")
	fmt.Println(
		"1- Igualdade\n2- Diferença\n3- Menor ou igual\n4- Menor\n5- Maior ou igual\n6- Maior",
	)
	_, err = fmt.Scan(&operacao)
	if err != nil {
		fmt.Println("Opção inválida")
		return
	}

	switch operacao {
	case 1:
		fmt.Printf("%v é igual a %v?\n", num1, num2)
		fmt.Println(num1 == num2)
	case 2:
		fmt.Printf("%v é diferente que %v?\n", num1, num2)
		fmt.Println(num1 != num2)
	case 3:
		fmt.Printf("%v é menor ou igual a %v?\n", num1, num2)
		fmt.Println(num1 <= num2)
	case 4:
		fmt.Printf("%v é menor a %v?\n", num1, num2)
		fmt.Println(num1 < num2)
	case 5:
		fmt.Printf("%v é maior ou igual a %v?\n", num1, num2)
		fmt.Println(num1 >= num2)
	case 6:
		fmt.Printf("%v é maior a %v?\n", num1, num2)
		fmt.Println(num1 > num2)
	default:
		fmt.Println("Entrada inválida")
	}
}
