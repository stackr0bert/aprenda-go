// Crie um loop utilizando a sintaxe: for condition {}
// utilize-o para demostrar os anos desde que você nasceu.

package main

import "fmt"

var ano int

func main() {
	
	fmt.Println("Digite o ano que você nasceu:")
	fmt.Scan(&ano)
	
	for ano := ano <= 2024; ano++ {
	fmt.Println(ano)

  }

	}
	

