package main

import "fmt"

func main() {
	qualquercoisa := map[int]string{
		123: "muito legal",
		98:  "menos legal um pouquinho",
		983: "esse é massa",
		18:  "idade de ir pra festa",
	} // ranges:
	fmt.Println(qualquercoisa)
	for key, value := range qualquercoisa {
		fmt.Println(key, value)
	}
	total := 0
	for key := range qualquercoisa {
		total += key
	}
	fmt.Println()
	fmt.Printf("A soma das keys é igual a %v", total)
	fmt.Println()

	// deletando elementos de um map
	fmt.Println("Deletando o item 123, ficamos:")
	delete(qualquercoisa, 123) // mapa, item
	fmt.Println((qualquercoisa))
}
