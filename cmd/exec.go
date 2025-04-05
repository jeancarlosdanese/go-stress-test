// cmd/exec.go

package cmd

import (
	"fmt"
	"time"

	"github.com/jeancarlosdanese/go-stress-test/internal/report"
	"github.com/jeancarlosdanese/go-stress-test/internal/tester"
	"github.com/spf13/cobra"
)

// Variáveis para armazenar os parâmetros do comando
var (
	url         string
	requests    int
	concurrency int
)

// execCmd representa o comando "exec" que executa o teste de carga
var execCmd = &cobra.Command{
	Use:   "exec",
	Short: "Executa o teste de carga",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Iniciando teste de carga:")
		fmt.Println("- URL:", url)
		fmt.Println("- Requisições:", requests)
		fmt.Println("- Concorrência:", concurrency)
		fmt.Println()

		start := time.Now()
		results := tester.RunTest(url, requests, concurrency)
		duration := time.Since(start)
		report.Generate(results, duration)
	},
}

func init() {
	// Flags com opções curtas
	execCmd.Flags().StringVarP(&url, "url", "u", "", "URL do serviço a ser testado (obrigatório)")
	execCmd.Flags().IntVarP(&requests, "requests", "r", 100, "Número total de requisições")
	execCmd.Flags().IntVarP(&concurrency, "concurrency", "c", 10, "Número de requisições simultâneas")
	execCmd.MarkFlagRequired("url")

	// Adiciona o comando "exec" ao comando raiz
	rootCmd.AddCommand(execCmd)
}
