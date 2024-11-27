// Crie uma variável de tipo string utilizando uma raw string literal. Demonstre-a.

package main

import "fmt"

var text string

func main() {
	text = `Olá, isso é uma string...
        
          literal ;D`
	fmt.Print(text)
}
