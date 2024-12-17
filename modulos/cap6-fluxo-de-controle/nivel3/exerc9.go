// CRIE UM PROGRAMA QUE UTILIZE A DECLARAÇÃO SWITCH, ONDE O SWITCH STATEMENT SEJA UMA VARIAVEL
// DO TIPO STRING E IDENTIFICADOR "jogoFavorito"

package main

import "fmt"

var jogoFavorito string

func main() {
	fmt.Println("Escolha um jogo 🎮:")
	fmt.Print("lol, cs 1.6 ,csgo OU --> dica💡 <--\n")
	fmt.Print("👉 ")
	fmt.Scan(&jogoFavorito)

	switch jogoFavorito {
	case "lol":
		fmt.Println("eca! 🤮🤮🤮")
	case "cs 1.6":
		fmt.Println("Parabéns, tu é 🫚")
	case "csgo":
		fmt.Println("que saudades😢")
	case "dica":
		fmt.Println("Tudo menos lol")

	default:
		fmt.Println("Escolha um jogo válido (isso não inclui o lol 🤮🤮🤮)")
	}
}
