// use var para declarar 3 variáveis. Elas devm ter package-level scope. Não atribua valores
// a estas variáveis. Utilize os seguites identificadores e tipos para estas variáveis:
// 1. Identificador "x" deverá ser do tipo int
// 2. Identificador "y" deverá ser do tipo string
// 3. Identificador "z" deverá ser do tipo bool

package main

import "fmt"

var (
	x int
	y string
	z bool
)

func main() {
	// TIPOS
	fmt.Println("TIPOS:")
	fmt.Printf("A variável x é do tipo %T\n", x)
	fmt.Printf("A variável y é do tipo %T\n", y)
	fmt.Printf("A variável z é do tipo %T\n", z)
	fmt.Println("VALORES:")
	fmt.Printf("%v\n%v\n%v\n", x, y, z)
}
