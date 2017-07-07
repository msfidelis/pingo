package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	exibeMenu()

	comando := leComando()

	switch comando {

	case 1:
		fmt.Println("Monitorando")
		for {
			iniciarMonitoramento()
		}

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

func iniciarMonitoramento() {

	site := "https://random-status-code.herokuapp.com/"

	//Slice - Abstração do Array com tamanho fixo
	sites := []string{
		"https://random-status-code.herokuapp.com/",
		"https://google.com/",
		"https://caelum.com.br",
		"https://nanoshots.com.br"}

	fmt.Println(sites)

	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "está funcionando normalmente", "- Status:", resp.StatusCode)
	} else {
		fmt.Println("Site:", site, "Não está respondendo", "- Status:", resp.StatusCode)
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
