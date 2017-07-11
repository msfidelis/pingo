package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {

	exibeMenu()

	comando := leComando()

	switch comando {

	case 1:
		fmt.Println("Monitorando")
		iniciarMonitoramento()

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

//Inicia o monitoramento dos sites alvo
func iniciarMonitoramento() {
	sites := leSitesDoArquivo()
	//Looping infinito com for
	for {

		for _, site := range sites { //For range
			testaSite(site)
		}

		time.Sleep(5 * time.Second)
	}

}

//Faz um request HTTP para o site alvo e mostra o resultado
func testaSite(site string) {

	//Request simples
	resp, err := http.Get(site)

	//Tratamento de erro nos requests
	if err != nil {
		fmt.Println(err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "está funcionando normalmente", "- Status:", resp.StatusCode)
	} else {
		fmt.Println("Site:", site, "Não está respondendo", "- Status:", resp.StatusCode)
	}
}

//Mostra as opções do Menu
func exibeMenu() {

	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("3 - Sair do Programa")

}

//Lê um input do teclado
func leComando() int {

	var comandoLido int
	fmt.Scan(&comandoLido)

	return comandoLido

}

//Retorna um slice de sites lidos de dentro de um arquivo
func leSitesDoArquivo() []string {

	var sites []string

	arquivo, err := os.Open("sites.txt")

	//Exemplo de tratamento de erro
	if err != nil {
		fmt.Println("Erro ao abrir o arquivo: ", err)
	}

	//Leitura de arquivos
	leitor := bufio.NewReader(arquivo)

	//Lendo todas as linhas de um arquivo
	for {

		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)

		sites = append(sites, linha)

		if err == io.EOF {
			break
		}

	}

	arquivo.Close()
	return sites
}
