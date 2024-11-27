// Utilizando iota, crie 4 constantes cujos valores sejam os próximos 4 anos. Demonstre estes valores.
package main

import "fmt"

const (
	a = iota + 2025
	b
	c
	d
)

func main() {
	fmt.Print(a, b, c, d)
}
