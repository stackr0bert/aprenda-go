// Escreva um programa que mostre um número decimal, binário e hexadecimal.
package main

import "fmt"

var num int

func main() {
	fmt.Print("Digite um número: ")
	_, err := fmt.Scan(&num)
	if err != nil {
		fmt.Println("Erro ao ler o número:", err)
		return
	}
	fmt.Printf("em decimal: %d\n", num)
	fmt.Printf("em binário: %b\n", num)
	fmt.Printf("em hexadecimal: %#x\n", num)
}
