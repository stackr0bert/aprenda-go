package main

import (
	"fmt"
)

func main() {
	x := 3
	if (x%2 == 0) || (x%3 == 0) {
		fmt.Println("é multiplo de 2 OU de 3")
	}

	if (x%2 == 0) && (x%3 == 0) {
		fmt.Println("É multiplo de dois E também de três")
	}

	if (x%2 == 0) && !(x%3 == 0) { // operador "!" de NEGAÇÃO
		fmt.Println("É multiplo de dois mas não de três")
	}

	if !(x%2 == 0) && (x%3 == 0) {
		fmt.Println("É multiplo de três mas não de dois")
	}
}

// "||" or
// "&&"and
