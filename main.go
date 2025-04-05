// main.go

package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/jeancarlosdanese/go-stress-test/report"
	"github.com/jeancarlosdanese/go-stress-test/tester"
)

func main() {
	// Flags da CLI
	url := flag.String("url", "", "URL do serviço a ser testado (obrigatório)")
	requests := flag.Int("requests", 100, "Número total de requisições")
	concurrency := flag.Int("concurrency", 10, "Número de requisições simultâneas")

	flag.Parse()

	// Validação dos parâmetros
	if *url == "" {
		fmt.Println("Erro: o parâmetro --url é obrigatório")
		flag.Usage()
		os.Exit(1)
	}

	// Exibe informações da execução
	fmt.Println("Iniciando teste de carga:")
	fmt.Println("- URL:", *url)
	fmt.Println("- Requisições:", *requests)
	fmt.Println("- Concorrência:", *concurrency)
	fmt.Println()

	// Inicia contagem de tempo
	start := time.Now()

	// Executa o teste
	results := tester.RunTest(*url, *requests, *concurrency)

	// Calcula tempo total
	duration := time.Since(start)

	// Gera relatório
	report.Generate(results, duration)
}
