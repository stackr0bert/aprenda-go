// Crie constantes tipadas e não tipadas. Demostre seus valores --> olhar a documentação
package main

import "fmt"

const (
	a         = 1
	b int16   = 10
	c         = 50.2
	d float32 = 11.2
)

func main() {
	fmt.Printf("%v %T\n", a, a)
	fmt.Printf("%v %T\n", b, b)
	fmt.Printf("%v %T\n", c, c)
	fmt.Printf("%v %T\n", d, d)
}
