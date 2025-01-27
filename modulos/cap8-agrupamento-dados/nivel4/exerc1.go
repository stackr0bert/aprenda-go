// Usando uma literal composta:
// - Crire um array que suporte 5 valores do tipo int
// - Atribua valores aos seus índices
// - Utilize range e demonstre os valores do array.
// - Utilizando format printing, demostre o tipo do array.
package main

import "fmt"

func main() {
	array := [5]int{1, 2, 3, 4, 5}
	for i := range array {
		fmt.Println(i)
	}
	fmt.Println()
	fmt.Printf("O tipo é %T", array)
}
