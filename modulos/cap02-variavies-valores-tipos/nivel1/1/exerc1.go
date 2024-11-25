// utilizando o operador curto de declaração, atribua estes valores às variáveis
// com identificadores "x", "y" e "z" -> 42, "James Bond" e true

// depois demostre os valores nestas variáveis, com: uma única declaração print e com multiplas declarações
package main

import "fmt"

func main() {
	x := 42
	y := "James Bond"
	z := true

	fmt.Println("Em uma única declaração: ")
	fmt.Printf("x: %v, y: %v, z: %v\n", x, y, z)
	fmt.Println("Em várias declarações")
	fmt.Printf("x: %v\n", x)
	fmt.Printf("y: %v\n", y)
	fmt.Printf("z: %v\n", z)
}

// 👌 Completo
