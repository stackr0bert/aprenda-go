// Crie um tipo. O tipo subjacente deve ser int.
// Crie uma variável para este tipo, com o identificador "x", utilizando a palavra-chave var.
// Na função main:
// 1. Demostre o valor da variável "x"
// 2. Demostre o tipo da variável "x"
// 3. Atribua 42 à variável "x" utilizando o operador ":="
// 4. Demostre o valor da variável "x"

package main

import "fmt"

type golang_type int

var x golang_type

func main() {
	fmt.Printf("O valor da variável x é %v, seu tipo é : %T\n", x, x)
	x := 42

	fmt.Printf("O valor da variável x é %v, seu tipo é : %T", x, x)
}
