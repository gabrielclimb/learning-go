package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitoramento = 1
const delay = 5

func main() {

	exibeIntroducao()
	registraLog("site-falso", true)
	for {
		exibeMenu()
		comando := leComando()
		// comando := 1
		switch comando {
		case 0:
			fmt.Println("Saindo do programa")
			os.Exit(0)
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo logs...")
			imprimeLogs()
		default:
			fmt.Println("Não conheço essa merda.")
			os.Exit(-1)
		}
	}
}

func exibeIntroducao() {
	nome := "Gabriel"
	idade := 34
	versao := 1.1
	fmt.Println("Olá sr.", nome, "sua idade é", idade)
	fmt.Println("Este programa está na versão", versao)
}

func exibeMenu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("O comando escolhido foi", comandoLido)
	return comandoLido
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	sites := leSitesDoArquivo()
	for i := 0; i < monitoramento; i++ {
		for i, site := range sites {
			fmt.Println("Testando site", i+1, ":", site)
			testaSite(site)
		}
		time.Sleep(delay * time.Second)
	}
}

func testaSite(site string) {
	resp, _ := http.Get(site)

	statusCodeSucess := 200

	if resp.StatusCode == statusCodeSucess {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
		registraLog(site, true)
	} else {
		fmt.Println("Site:", site,
			"está com problemas. Status Code:", resp.StatusCode)
		registraLog(site, false)
	}
	fmt.Println("")
}

func leSitesDoArquivo() []string {

	var sites []string

	arquivo, err := os.Open("sites.txt")
	// arquivo, err := ioutil.ReadFile("sites.txt")

	trataErro(err)

	leitor := bufio.NewReader(arquivo)

	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)
		fmt.Println(linha)

		sites = append(sites, linha)

		if err == io.EOF {
			break
		}
	}
	arquivo.Close()
	return sites
}

func trataErro(err error) {
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}
}

func registraLog(site string, status bool) {
	arquivo, err := os.OpenFile(
		"log.txt",
		os.O_RDWR|os.O_CREATE|os.O_APPEND,
		0755)

	if err != nil {
		fmt.Println("Erro ao registrar o log:", err)
	}

	writeLine := time.Now().Format("02/01/2006 15:04:05") + " - " + site + "- Status: " + strconv.FormatBool(status) + "\n"
	arquivo.WriteString(writeLine)

	fmt.Println(arquivo)

	arquivo.Close()
}

func imprimeLogs() {
	arquivo, err := ioutil.ReadFile("log.txt")
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println(string(arquivo))
}
