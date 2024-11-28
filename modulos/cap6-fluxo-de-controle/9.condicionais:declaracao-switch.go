package main

import "fmt"

func main() {
	x := 10

	switch {

	case x < 5:
		fmt.Println("chis é menor que cinco")
	case x == 5:
		fmt.Println("chis é igual a cinco")
	case x > 5:
		fmt.Println("chis é maior que cinco")

	}
	// outro caso seria

	quemtanoescritoriohoje := "fernanda"

	switch quemtanoescritoriohoje {
	case "zezinho":
		fmt.Println("Quem tá no escritório hoje é o zezinho ")

	case "juca":
		fmt.Println("Quem tá no escritório hoje é o marquinhos")

	case "fernanda":
		fmt.Println("Quem tá no escritório hoje é a fernanda")
		fallthrough // quando o fallthrough é executado, a comparação de baixo é pulada " case "maria":"
	case "maria": // e o fmt.println abaixo de maria é chamado
		fmt.Println("e sempre que a fernanda vem, a maria vem também")
	case "robert":
		fmt.Println("Quem tá no escritório é o Robert")
	default: // equivalente ao else
		fmt.Println("tá vazio ou a balada tava ótima, ou é feriado")
		// outro caso

	}
	switch quemtanoescritoriohoje {
	case "zezinho", "maria":
		fmt.Println("Hoje quem tá na firma é o time 1")

	case "fernanda", "robert":
		fmt.Println("Hoje quem tá na firma é o time 2")

	}

	// outro caso:

	y := -1
	switch {
	case (y == 4), (y == 8):
		fmt.Println("é 4 ou 8")
	case (y > 10):
		fmt.Println(" é maior que 10")
	default:
		fmt.Println("Escolha outro valor")

	}
}
