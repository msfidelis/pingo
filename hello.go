package main

import "fmt"
import "os"

func main() {

	exibeMenu()

	comando := leComando()

	switch comando {

	case 1:
		fmt.Println("Monitorando")
		break

	case 2:
		fmt.Println("Exibindo Logs")
		break

	case 3:
		fmt.Println("Saindo")
		os.Exit(0)
		break

	default:
		fmt.Println("Não conheço essa parada")
		os.Exit(-1)
		break
	}

}

func exibeMenu() {
	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("3 - Sair do Programa")
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)

	return comandoLido
}
