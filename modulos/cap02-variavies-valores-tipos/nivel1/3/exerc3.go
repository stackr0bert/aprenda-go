// utilizando a solução do exercício anterior:
// 1. Em package-level scope, atribua os seguintes valores às variáveis:
// a) para "x" atribua 42
// b) para "y" atribua "James Bond"
// c) para "z" atribua true
// 2. Na função main:
// a) Use fmt.Sprintf para atribuir todos esses valores a uma única variável. Faça essa atribuição de tipo
// string a uma variável de nome "s" utulizando o operador curto de declaração.
// b) Demostre a variável "s"

package main

import "fmt"

var (
	x int    = 42
	y string = "James Bond"
	z bool   = true
	s string
)

func main() {
	s := fmt.Sprintf("%v\t%v\t%v\t", x, y, z)
	fmt.Printf(s)
}
