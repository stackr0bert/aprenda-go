package main

import "fmt"

func main() {
	amigos := map[string]int{
		"alfredo": 5551234,
		"joana":   9996674,
	}
	fmt.Println(amigos)
	fmt.Println(amigos["joana"])
	amigos["gopher"] = 4444 // é assim que se adc valores no map
	fmt.Print(amigos, "\n\n")

	// comma ok idiom (não é usado só para maps)
	if será, ok := amigos["uahsdhuhasd"]; !ok {
		fmt.Println("Não tem!")
	} else {
		fmt.Println(será)
	}
}
